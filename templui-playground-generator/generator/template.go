package generator

var handlerTemplate = `// Code generated by "templui-playground-generator"; DO NOT EDIT.
package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/jfbus/templui-playground/playground/views/components"
{{- range .ImportPaths }}{{ if ne . "strconv" }}
	"{{.}}"
{{- end }}{{ end }}
)

var binder = &echo.DefaultBinder{}

func ViewComponent(c echo.Context) error {
	switch c.Param("component") {
{{- range .Components }}
		case "{{.Name}}":
			components.{{.Name}}Section().Render(c.Request().Context(), c.Response().Writer)
{{- end }}
	}
	return nil 
}

func UpdateComponent(c echo.Context) error {
	switch c.Param("component") {
{{- range .Components }}
		case "{{.Name}}":
			def := {{.Package}}.D{}
{{- range .Fields}}{{if .Default}}
			def.{{.Name}}={{.Default}}
{{- end}}{{end}}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			{{.Package}}.C(def).Render(c.Request().Context(), c.Response().Writer)
{{- end }}
	}
	return nil 
}

var _ style.D
var _ size.Size
var _ position.Position
`

var sidebarTemplate = `// Code generated by "templui-playground-generator"; DO NOT EDIT.
package views

import (
	"github.com/jfbus/templui/components/a"
	"github.com/jfbus/templui/components/icon"
)

templ Sidebar() {
<ul class="divide-y divide-solid">
{{- range .Components }}
	<li class="flex gap-2 p-4 items-center">
		@a.C(a.D{
			Text: "{{.Name}}",
			Attributes: templ.Attributes{
				"hx-get":   "/view/{{.Name}}",
				"hx-target": "#body",
			},
		})
		@icon.C(icon.ChevronRight)
	</li>
{{- end}}
</ul>
}
`

var componentTemplate = `// Code generated by "templui-playground-generator"; DO NOT EDIT.
package components

import (
{{- range .ImportPaths }}
	"{{.}}"
{{- end }}
)

templ {{.Name}}Form() {
{{range .Fields}}
{{- if eq .InputType "select" }}
	@selectfield.C(selectfield.D{
		Name: "{{.Name}}",
		Label: "{{.Name}}",
		Options: []option.D{
			{
				Label: "Select a value",
			},
{{- range .Values}}
			{
				Label: "{{.Package}}.{{.Name}}",
				Value: {{.Value}},
			},
{{- end }}
		},
	})
{{- end }}
{{- if eq .InputType "input" }}
	@input.C(input.D{
		Name: "{{.Name}}",
		Label: "{{.Name}}",
{{- if .Default}}
		Value: {{.Default}},
{{- end }}
	})
{{- end }}
{{- if eq .InputType "checkbox" }}
	@checkbox.C(checkbox.D{
		Name: "{{.Name}}",
		Label: "{{.Name}}",
		Value: "true",
	})
{{- end }}
{{- end }}
}

templ {{.Name}}Section() {
	@ComponentViewer(Component{
		ID:      "{{.Name}}",
		Package: "{{.Package}}",
		Form:    {{.Name}}Form(),
		Preview: {{.Package}}.C({{.Package}}.D{
{{- range .Fields}}{{if .Default}}
			{{.Name}}:{{.Default}},
{{- end}}{{end}}
		}),
	})
}
`
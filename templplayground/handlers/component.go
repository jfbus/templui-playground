// Code generated by "templplaygroundgenerator"; DO NOT EDIT.
package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/jfbus/templui-playground/templplayground/views/components"
	"github.com/jfbus/templui/components/loader"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/selectfield"
	"github.com/jfbus/templui/components/selectfield/option"
	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/input"
	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/position"
	"github.com/jfbus/templui/components/checkbox"
	"github.com/jfbus/templui/components/social"
	"github.com/jfbus/templui/components/checkboxgroup"
	"github.com/jfbus/templui/components/form/validation/message"
	"github.com/jfbus/templui/components/modal"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/a"
	"github.com/jfbus/templui/components/buttongroup"
	"github.com/jfbus/templui/components/radio"
	"github.com/jfbus/templui/components/table"
	"github.com/jfbus/templui/components/table/row"
	"github.com/jfbus/templui/components/accordion"
	"github.com/jfbus/templui/components/text"
	"github.com/jfbus/templui/components/accordion/element"
	"github.com/jfbus/templui/components/dropdown"
	"github.com/jfbus/templui/components/dropzone"
	"github.com/jfbus/templui/components/inline"
	"github.com/jfbus/templui/components/toast"
	"github.com/jfbus/templui/components/label"
	"github.com/jfbus/templui/components/radiogroup"
	"github.com/jfbus/templui/components/textarea"
)

var binder = &echo.DefaultBinder{}

func ViewComponent(c echo.Context) error {
	switch c.Param("component") {
		case "A":
			components.ASection().Render(c.Request().Context(), c.Response().Writer)
		case "Accordion":
			components.AccordionSection().Render(c.Request().Context(), c.Response().Writer)
		case "Button":
			components.ButtonSection().Render(c.Request().Context(), c.Response().Writer)
		case "Buttongroup":
			components.ButtongroupSection().Render(c.Request().Context(), c.Response().Writer)
		case "Checkbox":
			components.CheckboxSection().Render(c.Request().Context(), c.Response().Writer)
		case "Checkboxgroup":
			components.CheckboxgroupSection().Render(c.Request().Context(), c.Response().Writer)
		case "Dropdown":
			components.DropdownSection().Render(c.Request().Context(), c.Response().Writer)
		case "Dropzone":
			components.DropzoneSection().Render(c.Request().Context(), c.Response().Writer)
		case "Icon":
			components.IconSection().Render(c.Request().Context(), c.Response().Writer)
		case "Inline":
			components.InlineSection().Render(c.Request().Context(), c.Response().Writer)
		case "Input":
			components.InputSection().Render(c.Request().Context(), c.Response().Writer)
		case "Label":
			components.LabelSection().Render(c.Request().Context(), c.Response().Writer)
		case "Loader":
			components.LoaderSection().Render(c.Request().Context(), c.Response().Writer)
		case "Modal":
			components.ModalSection().Render(c.Request().Context(), c.Response().Writer)
		case "Radio":
			components.RadioSection().Render(c.Request().Context(), c.Response().Writer)
		case "Radiogroup":
			components.RadiogroupSection().Render(c.Request().Context(), c.Response().Writer)
		case "Selectfield":
			components.SelectfieldSection().Render(c.Request().Context(), c.Response().Writer)
		case "Social":
			components.SocialSection().Render(c.Request().Context(), c.Response().Writer)
		case "Table":
			components.TableSection().Render(c.Request().Context(), c.Response().Writer)
		case "Textarea":
			components.TextareaSection().Render(c.Request().Context(), c.Response().Writer)
		case "Toast":
			components.ToastSection().Render(c.Request().Context(), c.Response().Writer)
	}
	return nil 
}

func UpdateComponent(c echo.Context) error {
	switch c.Param("component") {
		case "A":
			def := a.D{}
			def.Href="https://github.com/jfbus/templui"
			def.Text="Text"
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			a.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Accordion":
			def := accordion.D{}
			def.ID="accordion"
			def.Children=[]element.D{{Title: "Section 1", Content:text.C("Content 1")},{Title: "Section 2", Content:text.C("Content 2")}}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			accordion.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Button":
			def := button.D{}
			def.ID="button"
			def.Label="Label"
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			button.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Buttongroup":
			def := buttongroup.D{}
			def.Buttons=[]button.D{{Label:"A"},{Label:"B"},{Label:"C"},{Label:"D"}}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			buttongroup.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Checkbox":
			def := checkbox.D{}
			def.ID="checkbox"
			def.Name="checkbox"
			def.Label="Label"
			def.Value="Value"
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			checkbox.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Checkboxgroup":
			def := checkboxgroup.D{}
			def.Name="checkboxgroup"
			def.Checkboxes=[]checkbox.D{{Name: "foo", Value: "1", Label: "Choice 1"},{Name: "foo", Value: "2", Label:"Choice 2"}}
			def.Message=&message.D{Message: "Validation message"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			checkboxgroup.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Dropdown":
			def := dropdown.D{}
			def.Button=button.D{Label:"Open dropdown"}
			def.Header=templ.Raw(`<div class="font-medium text-normal">Lorem ipsum</div><div class="text-sm text-gray-500">dolor sit amet</div>`)
			def.Links=[][]a.D{{{Href:"#", Text: "Section 1 link 1"},{Href:"#", Text: "Section 1 link 2"}},{{Href:"#", Text: "Section 2 link 1"}}}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			dropdown.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Dropzone":
			def := dropzone.D{}
			def.ID="dropzone"
			def.Name="dropzone"
			def.Message=&message.D{Message: "Validation message"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			dropzone.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Icon":
			def := icon.D{}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			icon.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Inline":
			def := inline.D{}
			def.Value="Value"
			def.Edit=input.C(input.D{Name:"edit"})
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			inline.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Input":
			def := input.D{}
			def.ID="input"
			def.Name="input"
			def.Label="Label"
			def.Value="Value"
			def.Placeholder="Placeholder"
			def.Message=&message.D{Message: "Validation message"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			input.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Label":
			def := label.D{}
			def.Label="Label"
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			label.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Loader":
			def := loader.D{}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			loader.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Modal":
			def := modal.D{}
			def.ID="modal"
			def.Title="Title"
			def.Content="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus."
			def.Close=&button.D{Label:"Cancel", Style: button.StyleOutline}
			def.Confirm=&button.D{Label:"OK"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			modal.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Radio":
			def := radio.D{}
			def.ID="radio"
			def.Name="radio"
			def.Label="Label"
			def.Value="Value"
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			radio.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Radiogroup":
			def := radiogroup.D{}
			def.Name="radiogroup"
			def.Radios=[]radio.D{{Name: "foo", Value: "1", Label: "Choice 1"},{Name: "foo", Value: "2", Label:"Choice 2"}}
			def.Message=&message.D{Message: "Validation message"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			radiogroup.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Selectfield":
			def := selectfield.D{}
			def.ID="selectfield"
			def.Name="selectfield"
			def.Label="Label"
			def.Options=[]option.D{{Label:"Select a value"},{Label:"Option 1"},{Label:"Option 2"}}
			def.Message=&message.D{Message: "Validation message"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			selectfield.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Social":
			def := social.D{}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			social.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Table":
			def := table.D{}
			def.Header=&row.D{Cells: []string{"Name", "Description", ""}}
			def.Rows=[]row.D{{Cells: []string{"Lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}},{Cells: []string{"Ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}},{Cells: []string{"Dolor", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}}}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			table.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Textarea":
			def := textarea.D{}
			def.ID="textarea"
			def.Name="textarea"
			def.Label="Label"
			def.Value="Value"
			def.Placeholder="Placeholder"
			def.Message=&message.D{Message: "Validation message"}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			textarea.C(def).Render(c.Request().Context(), c.Response().Writer)
		case "Toast":
			def := toast.D{}
			def.ContainerID="toasts"
			def.Content="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus."
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			toast.C(def).Render(c.Request().Context(), c.Response().Writer)
	}
	return nil 
}

var _ style.D
var _ size.Size
var _ position.Position

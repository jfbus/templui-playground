package generator

import "slices"

type Type interface {
	ImportPath() []string
	String() string
	Package() string
	Parent() Type
	ValueToString(value string) string
	BaseType() Type
}

type BaseType int

const (
	Unknown BaseType = iota
	Integer
	String
	Bool
	Any
	Struct
)

func (t BaseType) String() string {
	switch t {
	case Integer:
		return "int"
	case String:
		return "string"
	case Bool:
		return "bool"
	case Any:
		return "any"
	case Struct:
		return "struct"
	}
	return "???"
}

func (t BaseType) ImportPath() []string {
	if t == Integer {
		return []string{"strconv"}
	}
	return nil
}

func (t BaseType) Package() string {
	return ""
}

func (t BaseType) Parent() Type {
	return nil
}

func (t BaseType) BaseType() Type {
	return t
}

func (t BaseType) ValueToString(value string) string {
	switch t {
	case Integer:
		return "strconv.Itoa(int(" + value + "))"
	case String:
		return "string(" + value + ")"
	default:
		return value
	}
}

type UserType struct {
	Import     string
	Pkg        string
	Name       string
	Array      bool
	Pointer    bool
	ParentType Type
}

func (t UserType) String() string {
	if t.Name == "" && t.ParentType != nil {
		return t.ParentType.String()
	}
	if t.Pkg == "" {
		return t.Name
	}
	return t.Pkg + "." + t.Name
}

func (t UserType) ImportPath() []string {
	switch {
	case t.ParentType == nil && t.Import == "":
		return nil
	case t.ParentType != nil && t.Import == "":
		return t.ParentType.ImportPath()
	case t.ParentType == nil:
		return []string{t.Import}
	default:
		return append(t.ParentType.ImportPath(), t.Import)
	}
}

func (t UserType) Package() string {
	return t.Pkg
}

func (t UserType) Parent() Type {
	return t.ParentType
}

func (t UserType) BaseType() Type {
	if t.ParentType == nil {
		return Unknown
	}
	return t.ParentType.BaseType()
}

func (t UserType) ValueToString(value string) string {
	switch {
	case t.ParentType != nil:
		return t.ParentType.ValueToString(value)
	default:
		return String.ValueToString(value)
	}
}

type Choice struct {
	ImportPath string
	Package    string
	Name       string
	Value      string
}

func (c Choice) String() string {
	return c.Package + "." + c.Name
}

type Field struct {
	Name       string
	Type       Type
	Values     []Choice
	ImportPath []string
	Default    string
	Editable   bool
}

const (
	None     = ""
	Input    = "input"
	Checkbox = "checkbox"
	Select   = "select"
)

func (f *Field) InputType() string {
	switch {
	case !f.Editable:
		return None
	case len(f.Values) > 0:
		return Select
	case f.Type.BaseType() == Bool:
		return Checkbox
	case f.Type.BaseType() == String:
		return Input
	}
	return Input
}

func (f *Field) ImportPaths() []string {
	var ips []string
	if f.Editable {
		for _, ip := range f.Type.ImportPath() {
			if ip != "" && ip != "github.com/jfbus/templ-components/components/style" {
				ips = append(ips, ip)
			}
		}
	}
	for _, value := range f.Values {
		if value.ImportPath != "" {
			ips = append(ips, value.ImportPath)
		}
	}
	for _, ip := range f.ImportPath {
		if ip != "" && !slices.Contains(ips, ip) {
			ips = append(ips, ip)
		}
	}
	switch f.InputType() {
	case Select:
		ips = append(ips,
			"github.com/jfbus/templ-components/components/selectfield",
			"github.com/jfbus/templ-components/components/selectfield/option",
		)
	case Input:
		ips = append(ips,
			"github.com/jfbus/templ-components/components/input",
		)
	case Checkbox:
		ips = append(ips,
			"github.com/jfbus/templ-components/components/checkbox",
		)
	}

	return ips
}

func (f *Field) BinderFunc() string {
	switch f.Type.BaseType() {
	case Integer:
		return "Int"
	case String:
		return "String"
	case Bool:
		return "Bool"
	}
	return ""
}

type Definition struct {
	Fields []Field
}

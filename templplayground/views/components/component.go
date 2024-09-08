package components

import "github.com/a-h/templ"

type Component struct {
	ID     string
	Form   templ.Component
	Viewer templ.Component
}

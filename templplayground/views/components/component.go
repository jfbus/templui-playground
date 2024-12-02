package components

import (
	"github.com/a-h/templ"
)

type Component struct {
	ID      string
	Package string
	Form    templ.Component
	Preview templ.Component
}

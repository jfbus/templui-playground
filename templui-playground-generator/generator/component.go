package generator

import (
	"slices"
)

type Component struct {
	Name       string
	Package    string
	ImportPath string
	Fields     []Field
	Ignore     bool
}

func (c Component) ImportPaths() []string {
	ips := []string{c.ImportPath}
	for _, field := range c.Fields {
		for _, ip := range field.ImportPaths() {
			if !slices.Contains(ips, ip) {
				ips = append(ips, ip)
			}
		}
	}
	return ips
}

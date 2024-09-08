package generator

import (
	"go/ast"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
	"maps"
	"slices"
)

type Generator struct {
	definitions map[string]Definition
	components  map[string]Component
	choices     map[string][]Choice
	types       map[string]Type
	imports     map[string]string
}

func NewGenerator() *Generator {
	return &Generator{
		definitions: map[string]Definition{},
		components:  map[string]Component{},
		choices:     map[string][]Choice{},
		types:       map[string]Type{},
		imports:     map[string]string{},
	}
}
func (g *Generator) Parse() {
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests: true,
	}
	pkgs, err := packages.Load(cfg, "github.com/jfbus/templ-components/components/...")
	if err != nil {
		log.Fatal(err)
	}
	for _, pkg := range pkgs {
		// fmt.Printf("pkg %#v %v\n", pkg.Name, pkg.Types.Path())
		for _, file := range pkg.Syntax {
			g.parseFile(pkg, file)
		}
	}
}

func (g *Generator) parseType(pkg *packages.Package, decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		typ := spec.(*ast.TypeSpec)
		switch typ.Name.Name {
		case "D":
			if st, ok := typ.Type.(*ast.StructType); ok {
				def := g.definitions[pkg.Name]
				for _, field := range st.Fields.List {
					ft := parseType(pkg, field.Type)
					nf := Field{
						Name: field.Names[0].Name,
						Type: ft,
					}
					g.parseFieldComment(pkg, field.Doc, &nf)
					switch nf.Name {
					case "Label", "Value", "Placeholder":
						nf.Default = `"` + nf.Name + `"`
					case "ID":
						nf.Default = pkg.Name
						continue
					}
					def.Fields = append(def.Fields, nf)
				}
				g.definitions[pkg.Name] = def
				t := UserType{Import: pkg.PkgPath, Name: typ.Name.Name, Pkg: pkg.Name}
				g.types[t.String()] = t
			}
		default:
			tt := parseType(pkg, typ.Type)
			t := UserType{Import: pkg.PkgPath, Name: typ.Name.Name, Pkg: pkg.Name, ParentType: tt}
			g.types[t.String()] = t
		}
	}
}

func parseType(pkg *packages.Package, e ast.Expr) Type {
	switch sel := e.(type) {
	case *ast.SelectorExpr:
		return UserType{Pkg: sel.X.(*ast.Ident).Name, Name: sel.Sel.Name}
	case *ast.Ident:
		switch sel.Name {
		case "int":
			return Integer
		case "string":
			return String
		case "bool":
			return Bool
		case "any":
			return Any
		default:
			return UserType{Pkg: pkg.Name, Name: sel.Name}
		}
	case *ast.ArrayType:
		return UserType{Array: true, ParentType: parseType(pkg, sel.Elt)}
	case *ast.StarExpr:
		return UserType{Pointer: true, ParentType: parseType(pkg, sel.X)}
	default:
		return UserType{}
	}
}

func (g *Generator) parseFieldComment(pkg *packages.Package, doc *ast.CommentGroup, f *Field) {
	if doc == nil || len(doc.List) == 0 {
		return
	}
	for _, c := range doc.List {
		comment := c.Text
		switch {
		case strings.HasPrefix(comment, "//playground:values:"):
			parts := strings.Split(strings.TrimPrefix(comment, "//playground:values:"), ",")
			lst := make([]Choice, 0, len(parts))
			for _, part := range parts {
				subparts := strings.Split(part, ".")
				switch len(subparts) {
				case 1:
					lst = append(lst, Choice{
						Package: pkg.Name,
						Name:    part,
					})
				case 2:
					lst = append(lst, Choice{
						Package: subparts[0],
						Name:    subparts[1],
					})
				}
			}
			f.Values = lst
		case strings.HasPrefix(comment, "//playground:default:"):
			f.Default = strings.TrimPrefix(comment, "//playground:default:")
		case strings.HasPrefix(comment, "//playground:import:"):
			f.ImportPath = append(f.ImportPath, strings.TrimPrefix(comment, "//playground:import:"))
		}
	}
}

func (g *Generator) parseConst(pkg *packages.Package, decl *ast.GenDecl) {
	var tname Type
	for _, spec := range decl.Specs {
		typ := spec.(*ast.ValueSpec)
		if typ.Type != nil {
			tname = parseType(pkg, typ.Type)
		}
		if tname == nil {
			continue
		}
		g.choices[tname.String()] = append(g.choices[tname.String()], Choice{
			Package: pkg.Name,
			Name:    typ.Names[0].Name,
		})
	}
}

func (g *Generator) parseFile(pkg *packages.Package, file *ast.File) {
	if strings.Contains(strings.TrimPrefix(pkg.PkgPath, "github.com/jfbus/templ-components/components/"), "/") {
		return
	}
	g.imports[pkg.Name] = pkg.PkgPath
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.TYPE:
				g.parseType(pkg, decl)
			case token.CONST:
				g.parseConst(pkg, decl)
			}
		case *ast.FuncDecl:
			if decl.Name.Name != "C" {
				continue
			}
			g.components[pkg.Name] = Component{
				Name:       strings.ToTitle(pkg.Name[:1]) + pkg.Name[1:],
				Package:    pkg.Name,
				ImportPath: pkg.PkgPath,
			}
		}
	}
}

func (g *Generator) DefaultChoices(typ Type, pkg string) []Choice {
	lst := g.choices[typ.String()]
	chs := make([]Choice, 0, len(lst))
	for _, choice := range lst {
		if choice.Package == pkg {
			chs = append(chs, choice)
		}
	}
	if len(chs) == 0 {
		for _, choice := range lst {
			if choice.Package == typ.Package() {
				chs = append(chs, choice)
			}
		}
	}
	return chs
}

func (g *Generator) fixTypes() {
	for _, t := range g.types {
		if ut, ok := t.(UserType); ok {
			if put, ok := ut.ParentType.(UserType); ok && put.ParentType == nil {
				put.ParentType = g.types[put.String()]
			}
		}
	}
	for pkg := range g.definitions {
		for i, f := range g.definitions[pkg].Fields {
			if ut, ok := f.Type.(UserType); ok {
				ut.ParentType = g.types[ut.String()]
				if ut.ParentType == nil {
					ut.Import = g.imports[ut.Package()]
				}
				f.Type = ut
				g.definitions[pkg].Fields[i] = f
			}
		}
	}
}

func (g *Generator) fixValues() {
	for t := range g.choices {
		typ := g.types[t]
		for i := range g.choices[t] {
			g.choices[t][i].Value = typ.ValueToString(g.choices[t][i].String())
		}
	}
	for pkg := range g.definitions {
		for i, f := range g.definitions[pkg].Fields {
			for j, v := range g.definitions[pkg].Fields[i].Values {
				if v.Package != "" && v.Package != f.Type.Package() {
					g.definitions[pkg].Fields[i].Values[j].ImportPath = g.imports[v.Package]
				}
				g.definitions[pkg].Fields[i].Values[j].Value = f.Type.ValueToString(v.String())
			}
		}
	}
}

func (g *Generator) fixComponentFields() {
	for pkg, c := range g.components {
		var fields []Field
		for _, field := range g.definitions[pkg].Fields {
			if len(field.Values) == 0 {
				field.Values = g.DefaultChoices(field.Type, pkg)
			}
			field.Editable = field.Name == "Label" || field.Type.BaseType() == Bool || field.Type.BaseType() == String || len(field.Values) > 0
			fields = append(fields, field)
		}
		if len(fields) == 0 {
			delete(g.components, pkg)
		} else {
			c.Fields = fields
			g.components[pkg] = c
		}
	}
}

func (g *Generator) Generate() error {
	g.fixTypes()
	g.fixValues()
	g.fixComponentFields()
	{
		ht, err := template.New("handler").Parse(handlerTemplate)
		if err != nil {
			return err
		}
		fd, err := os.Create(filepath.Join("handlers/component.go"))
		if err != nil {
			return err
		}
		defer fd.Close()
		err = ht.Execute(fd, g.components)
		if err != nil {
			return err
		}
	}
	{
		mt, err := template.New("main").Parse(mainTemplate)
		if err != nil {
			return err
		}
		fd, err := os.Create(filepath.Join("views/main.templ"))
		if err != nil {
			return err
		}
		defer fd.Close()
		err = mt.Execute(fd, g.components)
		if err != nil {
			return err
		}
	}
	{
		ct, err := template.New("component").Parse(componentTemplate)
		if err != nil {
			return err
		}
		for _, pkg := range slices.Sorted(maps.Keys(g.components)) {
			c := g.components[pkg]
			fd, err := os.Create(filepath.Join("views/components", pkg+".templ"))
			if err != nil {
				return err
			}
			defer fd.Close()
			err = ct.Execute(fd, c)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

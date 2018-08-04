package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type visitor int

type Config struct {
	Package string
	Output  string
	Append  bool
}

var (
	fs = token.NewFileSet()

	// Conf - lazygen  main config
	Conf = Config{
		Output: "output",
		Append: false,
	}
)

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	switch t := node.(type) {
	case *ast.FuncDecl:

		fmt.Println("Func finded", t.Name)

		start := fs.Position(t.Pos()).Line
		end := fs.Position(t.End()).Line

		if t.Doc != nil {
			for _, comment := range t.Doc.List {
				fmt.Println("comment is ", comment.Text)
				ok, pos := FindValidFunction(fs, comment)
				if ok {
					fmt.Println("Function finded", t.Name, pos)

					ScanFunction(ReplaceConfig{
						Filename:    "file.go",
						Start:       start,
						End:         end,
						ReplaceType: "Note",
					})
				}
			}
		}
	}
	return v
}

func main() {

	fmt.Println("Generating code by lazygen")
	fmt.Println("Args is ", os.Args)
	var v visitor

	f, err := parser.ParseFile(fs, "file.go", nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	Conf.Package = f.Name.Name
	fmt.Println("Config is ", Conf)

	ast.Walk(v, f)

}

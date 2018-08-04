package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
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

		start := fs.Position(t.Pos()).Line
		end := fs.Position(t.End()).Line

		if t.Doc != nil {
			for _, comment := range t.Doc.List {
				fmt.Println("comment is ", comment.Text)
				finded, types := CheckCommentParams(comment.Text)
				fmt.Println("types is ", types)
				if !finded {
					log.Println("Cant find -type param....")
				}
				// @TODO refactor to reciver
				ok, pos := FindValidFunction(fs, comment)
				if ok {
					find, curVar, curParam := FindFunctionParams(t)
					fmt.Println("Finded", find, curVar, curParam)
					fmt.Println("Function finded", t.Name, pos)
					if find {

						ScanFunction(ReplaceConfig{
							Filename:    "file.go",
							Start:       start,
							End:         end,
							CurrentType: curParam,
							CurrentVar:  curVar,
							ReplaceType: types[0],
						})
						break
					}
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

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type visitor int

type Config struct {
	File    FileConf
	Package string
	Append  bool
}

var (
	fs = token.NewFileSet()

	// Conf - lazygen  main config
	Conf = Config{}
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
				finded, types := CheckCommentParams(comment.Text)
				fmt.Println("Finded type is ", types)
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
							File:        Conf.File,
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
			fmt.Println("---------------------")
		}
	}
	return v
}

func main() {

	files := WalkFiles(".")

	fmt.Println("Generating code by lazygen")
	// Parsing flags

	for _, file := range files {
		var v visitor
		Conf.File = file
		f, err := parser.ParseFile(fs, file.Filename, nil, parser.ParseComments)

		if err != nil {
			panic(err)
		}

		Conf.Package = f.Name.Name

		ast.Walk(v, f)
	}

}

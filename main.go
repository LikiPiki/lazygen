package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

type visitor int

var fs = token.NewFileSet()

func readFileLines(filename string, start int, end int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("err", err)
	}
	scanner := bufio.NewScanner(f)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if lineNumber >= start && lineNumber <= end {
			fmt.Println("- ", line)
		}
		lineNumber++
	}
}

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	switch t := node.(type) {
	case *ast.FuncDecl:

		fmt.Println("Func finded", t.Name)

		start := fs.Position(t.Pos()).Line
		end := fs.Position(t.End()).Line

		readFileLines(fs.Position(t.Pos()).Filename, start, end)

		for _, comment := range t.Doc.List {
			fmt.Println("comment is ", comment.Text)
		}
	}
	return v
}

func main() {

	fmt.Println("Args is ", os.Args)

	var v visitor

	f, err := parser.ParseFile(fs, "file.go", nil, parser.ParseComments)
	packageName := f.Name
	fmt.Println("package", packageName)

	ast.Walk(v, f)

	if err != nil {
		panic(err)
	}

}

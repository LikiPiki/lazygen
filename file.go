package main

import "fmt"

type Note struct {
	Name    string
	Content string
}

//go:generate lazygen -name=model
func (note Note) Test() {
	fmt.Println("note is ", note)
}

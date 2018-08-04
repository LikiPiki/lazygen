package main

import "fmt"

type Note struct {
	Name    string
	Content string
}

//go:generate lazygen -name=model
func (note Note) Test() {
	fmt.Println("note is ", note)
	//Coment is somethingNote
}

func (note Note) anoterfunc(test string) {
	fmt.Println("test here", note)
}

type bytes byte

//go:gerate kek -name=model
func (note Note) anoterfunc2(test string) {
	fmt.Println("test here", note)
}

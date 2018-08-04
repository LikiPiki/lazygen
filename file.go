package main

import "fmt"

type Note struct {
	Name    string
	Content string
}

//go:generate lazygen -type=Content
func (note Note) Test() {
	fmt.Println("note is ", note)
	//Coment is somethingNote
}

func (note Note) anoterfunc(test string) {
	fmt.Println("test here", note)
}

type bytes byte

func anoterfunc2(test string) {
	fmt.Println("test here")
}

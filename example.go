package main

import "fmt"

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

//lazygen -type=Cat
func (dog Dog) SayHello() {
	fmt.Println("Hello", dog.Name)
}

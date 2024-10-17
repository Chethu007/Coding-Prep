package main

import "fmt"

type MyInterface interface {
	Method1a()
	Method2a()
	Method3a()
}

type MyType struct{}

func (mt MyType) Method1a() {
	fmt.Println("Implementation of Method1")
}

// Uncommenting this will result in a compilation error
func (mt MyType) Method2a() {
	fmt.Println("Implementation of Method2")
}

// Uncommenting this will result in a compilation error
func (mt MyType) Method3a() {
	fmt.Println("Implementation of Method3")
}

func main() {
	var mt MyInterface
	mt = MyType{}
	mt.Method1a()
}

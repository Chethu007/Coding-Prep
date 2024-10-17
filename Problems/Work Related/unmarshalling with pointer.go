package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"name"`
	Age  *int   `xml:"age"` // Pointer to handle optional field
}

func main() {
	data := []byte(`<person><name>John</name><age>12</age></person>`)
	var p Person
	err := xml.Unmarshal(data, &p)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Person age: %+v\n", *p.Age)

	if p.Age == nil {
		fmt.Println("Age is missing")
	}
}

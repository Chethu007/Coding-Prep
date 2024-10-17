package main

import (
	"fmt"
)

// Define your struct
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Create an instance of the struct
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Marshal the struct to JSON
	//jsonData, err := json.Marshal(person)
	//if err != nil {
	//	fmt.Println("Error marshalling JSON:", err)
	//	return
	//}

	// Print the JSON in one line using fmt.Sprintf
	jsonString := fmt.Sprintf("%v", person)
	fmt.Println(jsonString)
}

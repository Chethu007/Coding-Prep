package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error // Declare an error variable

	// Explicitly calling err.Error() when err is nil
	fmt.Println("Using err.Error():", err.Error()) // Output: ""

	// Using err directly when err is nil
	//fmt.Println("Using err directly:", err) // Output: <nil>

	// Simulate a function that returns an error
	result, err := divide2(10, 0)
	if err != nil {
		// Explicitly calling err.Error() when err is non-nil
		fmt.Println("Error:", err.Error()) // Output: division by zero

		// Using err directly when err is non-nil
		fmt.Println("Using err directly:", err) // Output: division by zero
		return
	}
	fmt.Println("Result:", result)
}

func divide2(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero") // Create a new error with a message
	}
	return x / y, nil
}

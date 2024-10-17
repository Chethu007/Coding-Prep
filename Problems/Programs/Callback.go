package main

import "fmt"

// Callback function type
type Callback func(int) int

// Function that takes a callback function as an argument
func processCallback(num int, callback Callback) {
	result := callback(num) // Execute the callback function with the provided number
	fmt.Println("Result after callback:", result)
}

// Example callback functions
func square(num int) int {
	return num * num
}

func cube(num int) int {
	return num * num * num
}

func main() {
	number := 5

	// Using square function as a callback
	processCallback(number, square)

	// Using cube function as a callback
	processCallback(number, cube)
}

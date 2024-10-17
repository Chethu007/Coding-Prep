package main

import "fmt"

func main() {
	// Example byte array
	byteArray := []byte("This is a sample byte array")

	// Convert byte array to string
	str := string(byteArray)

	// Print the first 10 letters
	fmt.Println("First 10 letters:", str[:10])

	// Print the last 10 letters
	fmt.Println("Last 10 letters:", str[len(str)-10:])

	fmt.Println("First 10 letters:", byteArray[:10])

	// Print the last 10 letters
	fmt.Println("Last 10 letters:", byteArray[len(byteArray)-10:])
}

package main

import "fmt"

func main() {
	i := 1
	for i <= 20 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("i: %v\n", i)
		}
		i++
	}
}

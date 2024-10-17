package main

import "fmt"

func main() {
	// your code goes here
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		if n == 3 {
			fmt.Println("3 2 1")
			continue
		}
		temp := n
		for temp > 0 {
			fmt.Print(temp, " ")
			temp -= 2
		}
		if temp == 0 {
			temp = 1
		} else {
			temp = 2
		}
		for temp < n {
			fmt.Print(temp, " ")
			temp += 2
		}
		fmt.Println()
	}
}

package main

import "fmt"

func main() {
	count := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b
			s := fmt.Sprintf("%v", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}
	println(count)
}

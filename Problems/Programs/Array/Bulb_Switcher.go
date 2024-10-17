package main

import "fmt"

func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	track := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			track[i] = true
		} else {
			track[i] = false
		}
	}
	var ceil int
	if n%2 == 0 {
		ceil = n / 2
	} else {
		ceil = n/2 + 1
	}
	println("Ceil", ceil)
	for i := 3; i <= ceil; i++ {
		for j := i; j <= n; j += i {
			if track[j] {
				track[j] = false
			} else {
				track[j] = true
			}
		}
	}
	count := 0
	for i := 1; i <= ceil; i++ {
		if track[i] {
			count++
		}
		println(i, "->", track[i])
	}
	for i := ceil + 1; i <= n; i++ {
		if !track[i] {
			count++
		}
		println(i, "->", track[i])
	}
	return count
}

func main() {
	fmt.Println(bulbSwitch(12))
}

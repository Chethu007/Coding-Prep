package main

import "fmt"

func numberOfChild(n int, k int) int {
	if n == 2 {
		if k%2 == 0 {
			return 0
		} else {
			return 1
		}
	}
	if k < n {
		return k
	}
	n -= 1
	if k%n == 0 {
		if k%(2*n) == 0 {
			return 0
		}
		return n
	}
	dir := k / n
	ind := k % n
	if ind != 0 {
		dir++
	}
	if dir%2 == 0 {
		return n - ind
	} else {
		return ind
	}
}

func main() {
	fmt.Println(numberOfChild(5, 19))
}

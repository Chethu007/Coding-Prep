package main

import "fmt"

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n = n / i
			}
			println("i=", i)
			if i != 2 && i != 3 && i != 5 {
				return false
			}
		}
	}
	if n > 1 {
		println("n=", n)
		if n != 2 && n != 3 && n != 5 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("Res: %v", isUgly(6))
}

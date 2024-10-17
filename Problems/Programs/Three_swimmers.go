package main

import (
	"fmt"
)

func min3(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func funcValue(a, p int64) int64 {
	if p%a == 0 {
		return 0
	}
	return a - (p % a)
}

func main() {
	var t int64
	fmt.Scanln(&t)

	for t > 0 {
		var p, a, b, c int64
		fmt.Scanln(&p, &a, &b, &c)

		x := min3(funcValue(a, p), funcValue(b, p))
		result := min3(x, funcValue(c, p))

		fmt.Println(result)
		t--
	}
}

package main

import (
	"fmt"
)

func generateKey(num1 int, num2 int, num3 int) int {
	res, mul := 0, 1
	for num1 != 0 || num2 != 0 || num3 != 0 {
		x1 := num1 % 10
		x2 := num2 % 10
		minVal := min(x1, x2)
		x3 := num3 % 10
		minVal = min(minVal, x3)
		res += minVal * mul
		num1 /= 10
		num2 /= 10
		num3 /= 10
		mul *= 10
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Res: ", generateKey(987, 879, 798))
}

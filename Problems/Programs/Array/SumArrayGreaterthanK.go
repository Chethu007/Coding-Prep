package main

import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minOperations(k int) int {
	if k == 1 {
		return 0
	}
	res := math.MaxInt32
	for i := 1; i <= k/2; i++ {
		num := i
		count := i - 1
		temp := num
		println("Current number: ", num, "with count value: ", count)
		for num <= k {
			num += temp
			count++
			println("Duplicating.... Current number: ", num, "with count value: ", count)
		}
		res = min(res, count)
	}
	return res
}

func main() {
	println("Res: ", minOperations(11))
}

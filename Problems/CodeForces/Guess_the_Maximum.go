package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func guessMax(arr []int, n int) int {
	res := math.MaxInt32
	for i := 1; i < n; i++ {
		curMax := max(arr[i], arr[i-1])
		if res > curMax {
			res = curMax
		}
	}
	return res - 1
}

func main() {
	var t, n int
	fmt.Scan(&t)
	// Loop through each test case
	for t > 0 {
		// Read n
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		result := guessMax(arr, n)
		fmt.Println(result)
		print()
		t--
	}
}

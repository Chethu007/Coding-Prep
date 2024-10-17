package main

import (
	"fmt"
	"math"
)

func maximumSum(arr []int) int {
	n := len(arr)
	sumF, sumB, maxArr := 0, 0, math.MinInt32
	sumStart := make([]int, n)
	sumEnd := make([]int, n)
	for i := 0; i < n; i++ {
		sumF = max(arr[i], sumF+arr[i])
		sumB = max(arr[n-1-i], sumB+arr[n-1-i])
		maxArr = max(maxArr, max(sumF, sumB))
		sumStart[i] = sumF
		sumEnd[n-1-i] = sumB
	}
	fmt.Printf("sumStart: %v\n", sumStart)
	fmt.Printf("sumEnd: %v\n", sumEnd)
	for i := 1; i < n-1; i++ {
		maxArr = max(maxArr, sumStart[i-1]+sumEnd[i+1])
	}
	return maxArr
}

func main() {
	println(maximumSum([]int{1, -2, 0, 3}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

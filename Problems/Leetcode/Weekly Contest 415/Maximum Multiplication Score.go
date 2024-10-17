package main

import "math"

func maxScore(a []int, b []int) int64 {
	var maxVal int64
	n := len(b)
	dp1 := make([]int64, n)
	dp2 := make([]int64, n)
	dp3 := make([]int64, n)
	dp4 := make([]int64, n)
	for i := 0; i < n; i++ {
		dp1[i] = math.MinInt32
		dp2[i] = math.MinInt32
		dp3[i] = math.MinInt32
		dp4[i] = math.MinInt32
	}
	dp1[0] = int64(a[0] * b[0])
	for i := 1; i < n; i++ {
		dp1[i] = max(dp1[i-1], int64(a[0]*b[i]))
	}
	dp2[1] = dp1[0] + int64(a[1]*b[1])
	for i := 2; i < n; i++ {
		dp2[i] = max(dp2[i-1], dp1[i-1]+int64(a[1]*b[i]))
	}
	dp3[2] = dp2[1] + int64(a[2]*b[2])
	for i := 3; i < n; i++ {
		dp3[i] = max(dp3[i-1], dp2[i-1]+int64(a[2]*b[i]))
	}
	dp4[3] = dp3[2] + int64(a[3]*b[3])
	maxVal = dp4[3]
	for i := 4; i < n; i++ {
		dp4[i] = max(dp4[i-1], dp3[i-1]+int64(a[3]*b[i]))
		maxVal = dp4[i]
	}
	return maxVal
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []int{-1, 4, 5, -2}
	b := []int{-5, -1, -3, -2, -4}
	println(maxScore(a, b))
}

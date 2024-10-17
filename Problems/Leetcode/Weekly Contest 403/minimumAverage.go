package main

import (
	"fmt"
	"sort"
)

func min1(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	res := 2501.0
	for i := 0; i < len(nums)/2; i++ {
		res = min1(res, float64(nums[i]+nums[len(nums)-1-i])/2)
	}
	return res
}

func main() {
	fmt.Println(minimumAverage([]int{1, 9, 8, 3, 10, 5}))
}

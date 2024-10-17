package main

import (
	"fmt"
	"math"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maximumTotalCost(nums []int) int64 {
	var res int64 = math.MinInt64
	var temp int64 = 0
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}
	pre := make([]int64, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			temp += int64(nums[i])
		} else {
			temp -= int64(nums[i])
		}
		pre[i] = temp
	}
	fmt.Println(pre)
	for i := 0; i < n-1; i++ {
		flag := 0
		var temp int64 = 0
		for j := i + 1; j < n; j++ {
			if flag == 0 {
				temp += int64(nums[j])
				flag = 1
			} else {
				temp -= int64(nums[j])
				flag = 0
			}
		}
		println(temp)
		res = max(res, pre[i]+temp)
	}
	res = max(res, pre[n-1])
	return res
}

func main() {
	fmt.Println(maximumTotalCost([]int{1, -1}))
}

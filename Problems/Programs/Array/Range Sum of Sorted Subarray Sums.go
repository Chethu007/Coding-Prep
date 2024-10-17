package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	mod := 1000000007
	var res []int
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			res = append(res, sum%mod)
		}
	}
	sort.Ints(res)
	fmt.Printf("Res: %d\n", res)
	sum := 0
	for i := left - 1; i < right; i++ {
		sum = (sum + res[i]) % mod
	}
	return sum
}

func main() {
	println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 10))
}

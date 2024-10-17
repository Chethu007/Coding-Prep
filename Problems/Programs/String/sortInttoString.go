package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	n := len(nums)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i]+res[j] > res[j]+res[i]
	})
	result := ""
	for _, val := range res {
		result += val
	}
	return result
}

func largestNumber1(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		println(a, b)
		return a+b > b+a
	})

	result := ""
	for _, val := range nums {
		result += strconv.Itoa(val)
	}
	return result
}

func main() {
	arr := []int{111311, 1113}
	fmt.Printf("res: %v\n", largestNumber(arr))
}

package main

import (
	"fmt"
	"sort"
)

// Using two pointers - O(nlogn)
func twoSum1(nums []int, target int) []int {
	x := make([]int, len(nums))
	copy(x, nums)
	sort.Ints(nums)
	fmt.Printf("%v %v\n", nums, x)
	//println(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return result(x, []int{nums[i], nums[j]})
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func result(nums []int, cur []int) []int {
	res := []int{}
	for i, val := range nums {
		if val == cur[0] || val == cur[1] {
			res = append(res, i)
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	sort.Ints(nums)
	fmt.Printf("%v\n", nums)
	//println(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

func main() {
	nums := []int{11, 2, 6, 5, 42, 7, 8, 9}
	res := twoSum1(nums, 53)
	fmt.Printf("%v\n", res)
}

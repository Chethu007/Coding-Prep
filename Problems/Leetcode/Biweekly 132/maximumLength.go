package main

import "fmt"

func maximumLength(nums []int, k int) int {
	n := len(nums)
	var temp []int
	count := 0
	var solve func(i int, temp []int)
	solve = func(i int, temp []int) {
		if i == n {
			if check(temp, k) {
				if len(temp) > count {
					count = len(temp)
				}
			}
			return
		}
		solve(i+1, temp)
		temp = append(temp, nums[i])
		solve(i+1, temp)
	}
	solve(0, temp)
	return count
}

func check(temp []int, k int) bool {
	count := 0
	for i := 0; i < len(temp)-1; i++ {
		if temp[i] != temp[i+1] {
			count++
		}
	}
	if count > k {
		return false
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1}
	k := 0
	fmt.Println(maximumLength(nums, k))
}

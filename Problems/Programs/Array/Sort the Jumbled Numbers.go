package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sortJumbled(mapping []int, nums []int) []int {
	m := make(map[string]string)
	for i := range mapping {
		m[strconv.Itoa(i)] = strconv.Itoa(mapping[i])
	}
	var solve func(num int) int
	solve = func(num int) int {
		s := strconv.Itoa(num)
		var res string = ""
		for i := range s {
			res += m[string(s[i])]
		}
		resint, _ := strconv.Atoi(res)
		return resint
	}
	mp := make(map[int]int)
	for i := range nums {
		mp[nums[i]] = solve(nums[i])
		fmt.Printf("nums[i]->%d, mp[nums[i]]->%v", nums[i], mp[nums[i]])
		println()
	}
	sort.Slice(nums, func(i, j int) bool {
		return mp[nums[i]] < mp[nums[j]]
	})
	return nums
}

func main() {
	fmt.Printf("Res: %v", sortJumbled([]int{7, 9, 4, 1, 0, 3, 8, 6, 2, 5}, []int{47799, 19021, 162535, 454, 95, 51890378, 404}))
}

package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	fmt.Printf("Map: %v\n", m)
	var res, temp, temp2 []int
	for _, val := range m {
		temp = append(temp, val)
	}
	sort.Ints(temp)
	fmt.Printf("Temp: %v\n", temp)
	for i := len(temp) - 1; i >= 0 && k > 0; i-- {
		temp2 = append(temp2, temp[i])
		k--
	}
	fmt.Printf("Temp2: %v\n", temp2)
	for i := 0; i < len(temp2); i++ {
		for key, val := range m {
			if val == temp2[i] {
				res = append(res, key)
				delete(m, key)
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Printf("Res: %v", topKFrequent(nums, 2))
}

package main

import (
	"fmt"
	"strconv"
)

func almostEqual(n1, n2 int) bool {
	s1 := strconv.Itoa(n1)
	s2 := strconv.Itoa(n2)
	l1, l2 := len(s1), len(s2)
	if l1 < l2 {
		s1 = fmt.Sprintf("%0*s", l2, s1)
	} else if l1 > l2 {
		s2 = fmt.Sprintf("%0*s", l1, s2)
	}
	var diff []int
	for i := range s1 {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}
	if len(diff) == 0 {
		return true
	}
	if len(diff) == 2 {
		temp := []rune(s1)
		temp[diff[0]], temp[diff[1]] = temp[diff[1]], temp[diff[0]]
		if string(temp) == s2 {
			return true
		}
	}
	return false
}

func countPairs(nums []int) int {
	res, n := 0, len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if almostEqual(nums[i], nums[j]) {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("Result: %v\n", countPairs([]int{5, 12, 8, 5, 5, 1, 20, 3, 10, 10, 5, 5, 5, 5, 1}))
}

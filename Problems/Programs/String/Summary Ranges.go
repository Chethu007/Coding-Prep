package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var res []string
	n, i := len(nums), 0
	for i < n {
		start := nums[i]
		for i+1 < n && nums[i]+1 == nums[i+1] {
			i++
		}
		if start != nums[i] {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		} else {
			res = append(res, strconv.Itoa(start))
		}
		i++
	}
	return res
}

func main() {
	fmt.Printf("%v", summaryRanges([]int{1, 2, 3, 4, 6}))
}

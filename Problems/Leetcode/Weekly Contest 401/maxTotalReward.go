package main

import (
	"fmt"
	"sort"
)

func maxTotalReward(arr []int) int {
	sort.Ints(arr)
	maxReward := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		curReward := arr[i]
		for j := i + 1; j < n; j++ {
			if curReward < arr[j] {
				println("Current reward is:", curReward, arr[j], i)
				curReward += arr[j]
			}
		}
		if maxReward < curReward {
			maxReward = curReward
		}
	}
	return maxReward
}

func main() {
	fmt.Println(maxTotalReward([]int{2, 15, 14, 18}))
}

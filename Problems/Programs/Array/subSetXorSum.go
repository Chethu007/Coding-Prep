package main

import "fmt"

func subsetXORSum(nums []int) int {
	n := 1 << len(nums)
	sum := 0
	for i := 1; i < n; i++ {
		temp := 0
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) >= 1 {
				temp ^= nums[j]
			}
		}
		sum += temp
	}
	return sum
}

func main() {
	nums := []int{1, 3}
	fmt.Printf("Res: %v", subsetXORSum(nums))
}

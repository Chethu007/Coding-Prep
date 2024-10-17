package main

import "fmt"

func removeDuplicates(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); {
		m[nums[i]]++
		if m[nums[i]] > 2 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Printf("Array: %v", nums)
	return len(nums)
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 0, 0, 1, 2, 1, 1}
	println("Res: ", removeDuplicates(nums))
}

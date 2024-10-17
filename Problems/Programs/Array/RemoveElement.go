package main

import "fmt"

func removeElement(nums []int, val int) int {
	n, count := len(nums), 0
	j := n - 1
	for i := n - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			count++
		}
	}
	fmt.Printf("Array: %v", nums)
	return count
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println("Res: ", removeElement(nums, 2))
}

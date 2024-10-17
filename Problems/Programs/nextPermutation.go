package main

import (
	"fmt"
	"math"
)

func nextPermutation(nums []int) {
	index := -1
	n := len(nums)
	//Find the index
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}
	//If not found, it will be the first permutation
	println("Index", index)
	if index == -1 {
		reverse(nums, 0, n-1)
		return
	}

	//Find the smallest in the sequence which is greater than nums[index]
	nextsmallest := math.MaxInt32
	var n_index int
	for i := n - 1; i > index; i-- {
		if nums[i] > nums[index] {
			if nextsmallest > nums[i] {
				nextsmallest = nums[i]
				n_index = i

			}
		}
	}
	println("Nextsmallest", nums[n_index], " index:", n_index)
	//Swap them
	nums[index], nums[n_index] = nums[n_index], nums[index]
	fmt.Printf("Penultimate: %v\n", nums)
	//Reverse the sequence as they are in decreacing order
	reverse(nums, index+1, n-1)
}

func reverse(nums []int, j, k int) {
	println("j and k", j, " ", k)
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
	fmt.Printf("Reverse: %v", nums)
}

func main() {
	nums := []int{1, 5, 4, 3, 2, 0}
	fmt.Printf("Input: %v\n", nums)
	nextPermutation(nums)
	fmt.Printf("Next Permutation: %v", nums)
}

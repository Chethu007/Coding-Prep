package main

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	n := len(queries)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	fmt.Printf("res array: %v", res)
	var temp []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			temp = append(temp, i)
		}
	}
	for i := 0; i < n; i++ {
		if queries[i] <= len(temp) {
			res[i] = temp[queries[i]-1]
		}
	}
	return res
}

// nums = [1,3,1,7], queries = [1,3,2,4], x = 1
// Output: [0,-1,2,-1]
func main() {
	nums := []int{1, 3, 1, 7}
	q := []int{10}
	fmt.Printf("Res: %v", occurrencesOfElement(nums, q, 5))
}

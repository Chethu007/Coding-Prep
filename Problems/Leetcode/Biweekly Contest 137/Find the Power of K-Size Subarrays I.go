package main

import "fmt"

//	func resultsArray(nums []int, k int) []int {
//		if k == 1 {
//			return nums
//		}
//		n := len(nums)
//		var res []int
//
//		var solve func(n int) int
//		solve = func(n int) int {
//			n1 := n - k
//			sum := n*(n+1)/2 - (n1 * (n1 + 1) / 2)
//			return sum
//		}
//
//		sum := 0
//		for i := 0; i < k-1; i++ {
//			sum += nums[i]
//		}
//
//		for i := k - 1; i < n; i++ {
//			sum += nums[i]
//			if solve(nums[i]) == sum {
//				res = append(res, nums[i])
//			} else {
//				res = append(res, -1)
//			}
//			sum -= nums[i+1-k]
//		}
//		return res
//	}
func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	n := len(nums)
	var res []int

	var solve func(i int) int
	solve = func(idx int) int {
		count := 0
		for i := idx; i < idx+k-1; i++ {
			if nums[i]+1 == nums[i+1] {
				count++
			} else {
				break
			}
		}
		if count == k-1 {
			return nums[idx+k-1]
		}
		return -1
	}
	flag := 0

	for i := 0; i <= n-k; i++ {
		if flag == 1 {
			if nums[i+k-2]+1 == nums[i+k-1] {
				res = append(res, nums[i+k-1])
			} else {
				res = append(res, -1)
				flag = 0
			}
		} else {
			flag = 0
			temp := solve(i)
			if temp == -1 {
				res = append(res, -1)
			} else {
				res = append(res, temp)
				flag = 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{4, 4, 5, 1, 4, 1, 4, 5, 6, 1}
	fmt.Printf("Res: %v\n", resultsArray(nums, 3))
}

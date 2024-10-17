package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	var first func(l, r int) int
	var second func(l, r int) int
	first = func(l, r int) int {
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] < target {
				l = mid + 1
			} else if nums[mid] >= target {
				r = mid
			}
		}
		return l
	}
	second = func(l, r int) int {
		res := -1
		for l <= r {
			mid := l + (r-l)/2
			if nums[mid] <= target {
				res = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return res
	}
	l, r := 0, n-1
	left := first(l, r)
	right := second(l, r)
	println(left, "-", right)
	if left == -1 || right == -1 || (nums[left] != target && nums[right] != target) {
		return []int{-1, -1}
	}
	if nums[left] == target && nums[right] != target {
		return []int{left, left}
	} else if nums[left] != target && nums[right] == target {
		return []int{right, right}
	}
	return []int{left, right}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

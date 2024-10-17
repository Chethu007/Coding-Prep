package main

import (
	"fmt"
	"sort"
)

//func smallestDistancePair(nums []int, k int) int {
//	var res []int
//	mergeSort(nums, 0, len(nums)-1, &res)
//	sort.Ints(res)
//	return res[k-1]
//}
//
//func mergeSort(nums []int, low, high int, res *[]int) {
//	if low >= high {
//		return
//	}
//	mid := low + (high-low)/2
//	mergeSort(nums, low, mid, res)
//	mergeSort(nums, mid+1, high, res)
//	smallestDistance(nums, low, mid, high, res)
//	merge(nums, low, mid, high)
//}
//
//func smallestDistance(nums []int, low, mid, high int, res *[]int) {
//	right := mid + 1
//	for i := low; i <= mid; i++ {
//		for j := right; j <= high; j++ {
//			*res = append(*res, absDiff(nums[i], nums[j]))
//		}
//	}
//}
//
//func absDiff(i, j int) int {
//	if i > j {
//		return i - j
//	}
//	return j - i
//}
//
//func merge(nums []int, low, mid, high int) {
//	var temp []int
//	left, right := low, mid+1
//	for left <= mid && right <= high {
//		if nums[left] < nums[right] {
//			temp = append(temp, nums[left])
//			left++
//		} else {
//			temp = append(temp, nums[right])
//			right++
//		}
//	}
//	for left <= mid {
//		temp = append(temp, nums[left])
//		left++
//	}
//	for right <= high {
//		temp = append(temp, nums[right])
//		right++
//	}
//	for i := low; i <= high; i++ {
//		nums[i] = temp[i-low]
//	}
//}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	var absDiff func(i, j int) int
	absDiff = func(i, j int) int {
		if i > j {
			return nums[i] - nums[j]
		}
		return nums[j] - nums[i]
	}
	maxVal := -1
	for i := range nums {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	low, high := 0, maxVal
	for low < high {
		mid := low + (high-low)/2
		left, right, count := 0, 0, 0
		for right < n {
			for absDiff(left, right) > mid {
				left++
			}
			count += right - left
			right++
		}
		if count >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}

func main() {
	fmt.Printf("Res: %v\n", smallestDistancePair([]int{1, 6, 1}, 3))
}

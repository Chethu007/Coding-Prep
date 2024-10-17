package main

import "fmt"

func merge(nums []int, low, mid, high int) {
	var temp []int
	left, right := low, mid+1
	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, nums[left])
		left++
	}
	for right <= high {
		temp = append(temp, nums[right])
		right++
	}
	for i := low; i <= high; i++ {
		nums[i] = temp[i-low]
	}
}

func countReverse(nums []int, low, mid, high int) int {
	count, right := 0, mid+1
	for i := low; i <= mid; i++ {
		for right <= high && nums[i] > 2*nums[right] {
			right++
		}
		count += right - (mid + 1)
	}
	return count
}

func mergeSort(nums []int, low, high int) int {
	count := 0
	if low >= high {
		return 0
	}
	mid := low + (high-low)/2
	count += mergeSort(nums, low, mid)
	count += mergeSort(nums, mid+1, high)
	count += countReverse(nums, low, mid, high)
	merge(nums, low, mid, high)
	return count
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{1, 3, 2, 3, 1}
	println(reversePairs(nums))
	fmt.Println(nums)
}

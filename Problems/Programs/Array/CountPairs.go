package main

import "fmt"

func Merge(nums []int, low, mid, high int) int {
	count := 0
	var temp []int
	left, right := low, mid+1
	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			count += mid - left + 1
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
	return count
}

func mSort(nums []int, low, high int) int {
	count := 0
	if low >= high {
		return 0
	}
	mid := low + (high-low)/2
	count += mSort(nums, low, mid)
	count += mSort(nums, mid+1, high)
	count += Merge(nums, low, mid, high)
	return count
}

func inversionCount(nums []int) int {
	return mSort(nums, 0, len(nums)-1)
}

func subsets(nums []int) [][]int {
	// Bit Manipulation
	// n := len(nums)
	// subset := 1<<n
	// var res [][]int
	// for i:=0; i<subset; i++ {
	//     temp := []int{}
	//     for j:=0; j<n; j++{
	//         if i&(1<<j) >= 1{
	//             temp = append(temp, nums[j])
	//         }
	//     }
	//     res = append(res, temp)
	// }
	// return res

	// Recursion
	// var res [][]int
	// n := len(nums)
	// var temp []int
	// f(0,n,temp,nums,res)
	// return res
	var res [][]int
	var temp []int
	f(0, nums, temp, &res)
	return res
}

func f(index int, nums []int, temp []int, res *[][]int) {
	if index == len(nums) {
		// Append a copy of temp to res
		// This is done to ensure that the contents of temp at the current point in the recursion are preserved separately.
		// Preventing Mutability Issues and Preserving Different States
		copyTemp := make([]int, len(temp))
		copy(copyTemp, temp)
		*res = append(*res, copyTemp) // Update the value pointed to by res
		return
	}

	// Include current element

	f(index+1, nums, temp, res)
	temp = append(temp, nums[index])
	// Exclude current element
	// temp = (temp)[:len(temp)-1]
	f(index+1, nums, temp, res)
}

func main() {
	nums := []int{12, 11, 10}
	println(inversionCount(nums))
	fmt.Println(nums)
}

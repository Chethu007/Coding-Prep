package main

import (
	"fmt"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums1)
	fmt.Printf("Nums1: %v", nums1)
	fmt.Printf("Nums2: %v\n", nums2)
	mini := 1001
	for i := 0; i <= 2; i++ {
		k := 0
		diff := nums2[0] - nums1[i]
		skip := i
		for j := i; j < len(nums1); j++ {
			if nums2[k]-nums1[j] != diff && skip == 2 {
				break
			}
			println("I value:", i, "J value:", j, "nums2[k]-nums1[j] = ", nums2[k], "-", nums1[j], " = Diff: ", diff)
			if nums2[k]-nums1[j] != diff {
				skip++
			} else {
				k++
			}
			if k == len(nums2) {
				break
			}
		}
		println("K value:", k)
		if k == len(nums2) {
			println("Diff:", diff)
			if mini < 0 && diff < 0 {
				mini = diff
			} else {
				mini = min2(mini, diff)
			}
		}
	}
	if mini == 1001 {
		return 0
	}
	return mini
}

func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//	nums1 = [4,20,16,12,8], nums2 = [14,18,10]
	nums1 := []int{4, 20, 16, 12, 8}
	nums2 := []int{14, 18, 10}
	println("Res:", minimumAddedInteger(nums1, nums2))
}

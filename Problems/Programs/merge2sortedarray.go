package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, z := 0, 0, 0
	arr := make([]int, m+n)
	for i < m && j < n {
		//if nums1[i]==0 {
		//	i++
		//	continue
		//}
		if nums1[i] < nums2[j] {
			arr[z] = nums1[i]
			z++
			i++
		} else {
			arr[z] = nums2[j]
			z++
			j++
		}
	}
	for i < m {
		arr[z] = nums1[i]
		z++
		i++
	}
	for j < n {
		arr[z] = nums2[j]
		z++
		j++
	}
	//for i := 0; i < m+n; i++ {
	//	nums1[i] = arr[i]
	//}
	nums1 = append([]int(nil), arr...)
	return nums1
}

func merge1(nums1 []int, m int, nums2 []int, n int) []int {
	for m > 0 && n > 0 {
		println("Value of m, n:", m, " ", n)
		if nums1[m-1] <= nums2[n-1] {
			nums1[m+n-1] = nums2[m-1]
			m--
		} else {
			nums1[m+n-1] = nums1[n-1]
			n--
		}
	}

	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}

	return nums1
}

func main() {
	arr1 := []int{1, 0}
	arr2 := []int{1, 0}
	//fmt.Printf("Result %v\n", merge(arr1, 9, arr2, 5))
	fmt.Printf("Result %v\n", merge1(arr1, 1, arr2, 1))
}

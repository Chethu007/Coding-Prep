package main

import "fmt"

//	func isArraySpecial(nums []int, queries [][]int) []bool {
//		n := len(queries)
//		var res []bool
//		for i := 0; i < n; i++ {
//			res = append(res, isArraySpecial1(nums, queries[i][0], queries[i][1]))
//			println("Queries", queries[i][0], queries[i][1])
//		}
//		return res
//	}
//
//	func isArraySpecial1(nums []int, start, end int) bool {
//		if nums[start]%2 == 0 {
//			for i := start + 1; i <= end; i++ {
//				if (start+i)%2 == 0 {
//					if nums[i]%2 != 0 {
//						return false
//					}
//				} else if nums[i]%2 != 1 {
//					return false
//				}
//			}
//		} else {
//			for i := start + 1; i <= end; i++ {
//				if (start+i)%2 == 0 {
//					if nums[i]%2 != 1 {
//						return false
//					}
//				} else if nums[i]%2 != 0 {
//					return false
//				}
//			}
//		}
//		return true
//	}
//func isArraySpecial(nums []int, queries [][]int) []bool {
//	n1 := len(queries)
//	m := make(map[int]bool)
//	n := len(nums)
//	m[n-1] = true
//	for i := 0; i < n-1; i++ {
//		if (nums[i]%2 == 0 && nums[i+1]%2 == 1) || (nums[i]%2 == 1 && nums[i+1]%2 == 0) {
//			m[i] = true
//		} else {
//			m[i] = false
//		}
//	}
//	//fmt.Printf("Map: %v", m)
//	var res []bool
//	for i := 0; i < n1; i++ {
//		flag := true
//		for j := queries[i][0]; j < queries[i][1]; j++ {
//			if m[j] == false {
//				flag = false
//				break
//			}
//		}
//		res = append(res, flag)
//	}
//	return res
//}

func isArraySpecial(nums []int, queries [][]int) []bool {
	n, n1, count := len(nums), len(queries), 0
	violations := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i]%2 == nums[i-1]%2 {
			count++
		}
		violations[i] = count
	}
	fmt.Printf("Violaton arr: %v\n", violations)
	var res []bool
	for i := 0; i < n1; i++ {
		start := queries[i][0]
		end := queries[i][1]
		res = append(res, violations[end]-violations[start] == 0)
	}
	return res
}

func main() {
	nums := []int{4, 1, 3, 6}
	queries := [][]int{{0, 2}, {2, 3}}
	fmt.Printf("Res: %v", isArraySpecial(nums, queries))
}

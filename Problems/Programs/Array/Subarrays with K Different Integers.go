package main

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	n := len(nums)
	if k == 1 {
		return n
	}
	m := make(map[int]int)
	res := 0
	l, r := 0, 0
	for r < n {
		if len(m) < k {
			m[nums[r]]++
			if len(m) == k {
				res++
			}
			r++
			//println("r1 ", r, "len of map ", len(m))
		} else if len(m) == k {
			//println("r2 ", r, "len of map ", len(m))
			//println("RES1: ", res)
			for r < n && m[nums[r]] != 0 {
				res++
				m[nums[r]]++
				r++
				//println("RES2: ", res, "nums[r+1]", nums[r+1])
			}
			for len(m) == k {
				m[nums[l]]--
				if m[nums[l]] == 0 {
					delete(m, nums[l])
				}
				if len(m) == k {
					//println("RES3: ", res)
					res++
				}
				l++
			}
		}
	}
	for len(m) == k {
		m[nums[l]]--
		if m[nums[l]] == 0 {
			delete(m, nums[l])
		}
		if len(m) == k {
			//println("RES3: ", res)
			res++
		}
		l++
	}
	return res
}

func main() {
	fmt.Printf("Result: %d\n", subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}

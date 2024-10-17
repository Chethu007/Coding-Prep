package main

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	total, l, r, winLen, n := nums[0], 0, 0, 1, len(nums)
	maxlen := 1
	for r < n && l < n {
		if r+1 < n && nums[r]*winLen <= total+k {
			println("Withing window: nums[r] is ", nums[r], " window length ", winLen, " and r is", r)
			println("total+k: ", total+k)
			total += nums[r+1]
			r++
			maxlen = max(maxlen, winLen)
			winLen++
		} else {
			println("Reduce window: nums[l] is ", nums[l], " window length ", winLen, " and l is", l)
			println("total+k: ", total+k)
			total -= nums[l]
			l++
			maxlen = max(maxlen, winLen)
			winLen--
		}
	}
	maxlen = max(maxlen, winLen)
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 4}
	println("Res: ", maxFrequency(nums, 2))
}

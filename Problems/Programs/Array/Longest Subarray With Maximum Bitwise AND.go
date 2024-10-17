package main

func longestSubarray(nums []int) int {
	maxVal, maxSubArray, n := 0, 0, len(nums)
	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] == maxVal {
			j := i
			for j+1 < n && nums[j] == nums[j+1] {
				j++
			}
			maxSubArray = max4(maxSubArray, j-i+1)
			i = j
		}
	}
	return maxVal
}

func max4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(longestSubarray([]int{1, 2, 3, 4, 4, 4}))
}

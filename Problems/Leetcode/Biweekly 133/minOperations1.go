package main

func minOperations1(nums []int) int {
	count, n := 0, len(nums)
	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			for j := i; j < i+3; j++ {
				if nums[j] == 0 {
					nums[j] = 1
				} else {
					nums[j] = 0
				}
			}
			count++
		}
	}
	for _, val := range nums {
		if val == 0 {
			return -1
		}
	}
	return count
}

func main() {
	println(minOperations1([]int{0, 1, 1, 1, 0, 0}))
}

package main

func findMaximumScore(nums []int) int64 {
	n := len(nums)
	var res int64 = 0
	i, j := 0, 1
	for j < n {
		for j < len(nums)-1 && nums[i] > nums[j] {
			j++
		}
		res += int64((j - i) * (nums[i]))
		i = j
		j++
	}
	return res
}

func main() {
	println(findMaximumScore([]int{2, 3, 4, 2, 5, 6, 7, 2, 8}))
}

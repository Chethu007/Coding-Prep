package main

func minSwaps(nums []int) int {
	totalCount, n := 0, len(nums)
	for i := range nums {
		if nums[i] == 1 {
			totalCount++
		}
	}
	if n == totalCount {
		return 0
	}
	res, count := totalCount, 0
	for i := 0; i < totalCount; i++ {
		if nums[i] == 1 {
			count++
		}
		//println(count)
	}
	res = min12(res, totalCount-count)
	l, r := 1, totalCount
	for l < n {
		if nums[(l-1)%n] == 1 {
			count--
		}
		if nums[r%n] == 1 {
			count++
		}
		//println(l, "-->", count)
		res = min12(res, totalCount-count)
		l++
		r++
	}
	return res
}

func main() {
	println(minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))
}

func min12(a, b int) int {
	if a < b {
		return a
	}
	return b
}

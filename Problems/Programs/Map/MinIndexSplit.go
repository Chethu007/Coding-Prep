package main

func minimumIndex(nums []int) int {
	m := make(map[int]int)
	dominant, dominantCount := 0, 0
	for _, val := range nums {
		m[val]++
		if dominantCount < m[val] {
			dominantCount = m[val]
			dominant = val
		}
	}
	println("Dominant element: ", dominant)
	if m[dominant] == 1 {
		return -1
	}
	n := len(nums)
	validIndex := -1
	for i := 0; i < n-1; i++ {
		x, y := 0, 0
		for j := 0; j <= i; j++ {
			if nums[j] == dominant {
				x++
			}
		}
		for k := i + 1; k < n; k++ {
			if nums[k] == dominant {
				y++
			}
		}
		println("X: ", x, " Y: ", y)
		if x > (i+1)/2 && y > (n-i-1)/2 {
			validIndex = i
			break
		}
	}
	return validIndex
}

func main() {
	nums := []int{3, 1, 1, 1}
	println("Res: ", minimumIndex(nums))
}

package main

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(r, b int) int {
	count, flag, i := 0, 1, 1
	println("Red and Blue", r, b)
	for r >= 0 && b >= 0 {
		if flag == 1 {
			r -= i
			if r < 0 {
				println("Exit 1", count)
				return count
			}
			i++
			flag = 0
		} else {
			b -= i
			if b < 0 {
				println("Exit 2", count)
				return count
			}
			i++
			flag = 1
		}
		count++
	}
	println("Exit 3", count)
	return count
}

func maxHeightOfTriangle(red int, blue int) int {
	return max3(solve(red, blue), solve(blue, red))
}

func maximumLength(nums []int) int {
	res, i, n := 0, 0, len(nums)
	for i < n-1 {
		count := 1
		for i < n-1 && nums[i]%2 == nums[i+1]%2 {
			i++
			count++
		}
		if res < count {
			res = count
		}
		i++
	}
	i = 0
	for i < n-1 {
		count := 1
		for i < n-1 && nums[i]%2 != nums[i+1]%2 {
			i++
			count++
		}
		if res < count {
			res = count
		}
		i++
	}
	return res
}

func main() {
	//println(maxHeightOfTriangle(2, 4))
	println(maximumLength([]int{1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 8, 8, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8}))
}

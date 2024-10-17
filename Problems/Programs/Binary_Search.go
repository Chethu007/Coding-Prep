package main

func binartsearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// Return directly
			return mid
		}
	}
	// Return directly
	return -1
}

func leftbound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		print("left, mid, right:", left, " ", mid, " ", right, " ")
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	println()
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightbound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		print("left, mid, right:", left, " ", mid, " ", right, " ")
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	println()
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func Upper(s rune) rune {
	if s >= 'a' && s <= 'z' {
		return s - 'a' + 'A'
	}
	return s
}

func Lower(s rune) rune {
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return s
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 7, 8, 9}
	target := 4
	println("Binary search res", binartsearch(nums, target))
	println("Left bound res", leftbound(nums, target))
	println("Right bound res", rightbound(nums, target))
	//var x, y int
	//fmt.Scanf("%d %d", &x, &y)
	//if (x ^ y) < 0 {
	//	println("Signs are different")
	//} else {
	//	println("Signs are same")
	//}
	str := "Hello World!!!!"
	//println("Upper case:", strings.ToUpper(str))
	//println("Lower case:", strings.ToLower(str))
	con := make([]rune, len(str))
	for i, ch := range str {
		con[i] = Upper(ch)
		println(con[i])
	}
	println("Upper case:", string(con))
	con1 := make([]rune, len(str))
	for i, ch := range str {
		con1[i] = Lower(ch)
	}
	println("Lower case:", string(con1))
}

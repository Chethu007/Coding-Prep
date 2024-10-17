package main

func partitionDisjoint(nums []int) int {
	n, maxNum, maxlen := len(nums), -1, 0
	var jVal int
	for i := 0; i < n-1; i++ {
		println("Partition index i: ", i, " and the element: ", nums[i])
		for j := i + 1; j < n; j++ {
			jVal = j
			if nums[i] > nums[j] {
				if maxNum < nums[i] {
					maxNum = nums[i]
				}
				println("MaxNum: ", maxNum)
				break
			}
			println("jVal: ", jVal)
		}
		if jVal == n-1 && nums[i] > maxNum {
			println("Maxlen: ", maxlen)
			maxlen = Max(maxlen, i)
		}
	}
	return maxlen
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	nums := []int{5, 0, 3, 8, 6}
	println("Res: ", partitionDisjoint(nums))
}

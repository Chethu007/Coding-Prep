package main

func mostFrequentEven(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	maxCount := -1
	res := 100001
	for key, val := range m {
		if key%2 == 0 {
			println("Key: ", key, " Value: ", val)
			if maxCount <= val {
				println("MaxCount: ", maxCount)
				maxCount = val
			}
		}
	}
	if maxCount == -1 {
		return -1
	}
	for k, v := range m {
		if v == maxCount && k%2 == 0 {
			if res > k {
				res = k
			}
		}
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 2, 4, 4, 1}
	println("Res: ", mostFrequentEven(nums))
}

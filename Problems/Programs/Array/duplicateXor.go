package main

func duplicateNumbersXOR(nums []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			res ^= nums[i]
		} else {
			m[nums[i]]++
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 2, 1}
	println("Res: ", duplicateNumbersXOR(nums))
}

package main

func minimumOperations(nums []int) int {
	count := 0
	for _, val := range nums {
		val = val % 3
		if val > 3-val {
			val = 3 - val
		}
		count += val
	}
	return count
}

func main() {
	println(minimumOperations([]int{3, 6, 9}))
}

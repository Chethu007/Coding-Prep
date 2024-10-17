package main

func main() {
	nums := []int{4, 16, 23, 65, 5, 3, 1}
	max := nums[0]
	for _, num := range nums[1:] {
		if max < num {
			max = num
		}
	}
	println(max)
}

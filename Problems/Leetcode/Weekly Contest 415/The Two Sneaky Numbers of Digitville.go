package main

import "fmt"

func getSneakyNumbers(nums []int) []int {
	var res []int
	m := make(map[int]int)
	for _, val := range nums {
		if m[val] == 0 {
			m[val] = 1
		} else {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	fmt.Printf("Res: %v", getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}

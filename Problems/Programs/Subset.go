package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	subset := (1 << n)
	res := [][]int{}
	for i := 0; i < subset; i++ {
		println("Value of i", i)
		temp := []int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) >= 1 {
				fmt.Printf("i: %b, (1<<j): %b", i, 1<<j)
				println(" Value of n", n, " value of nums[j]", nums[j])
				temp = append(temp, nums[j])
			}
		}
		fmt.Printf("Temp: %v\n", temp)
		res = append(res, temp)
	}
	return res
}

func main() {
	sub := []int{1, 2, 3}
	x := subsets(sub)
	fmt.Printf("Res: %v\n", x)
	a := 7 & 3
	fmt.Printf("Res: %v\n", a)
}

package main

import "fmt"

func stableMountains(height []int, threshold int) []int {
	var res []int
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Printf("Res: %v", stableMountains([]int{10, 1, 10, 1, 10}, 10))
}

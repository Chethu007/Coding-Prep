package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	// Brute Force Solution
	// prod,n:=math.MinInt32,len(nums)
	// for i:=0; i<n; i++ {
	//     res := 1
	//     for j:=i; j<n; j++ {
	//         res *= nums[j]
	//         prod = max(prod, res)
	//     }
	// }
	// return prod

	// Optimal Solution based on the observation of postive, negative and 0 elements in the array
	n := len(nums)
	maxProd := math.MinInt32
	premul, suffmul := 1, 1
	for i := 0; i < n; i++ {
		premul *= nums[i]
		suffmul *= nums[n-i-1]
		maxProd = max(maxProd, max(premul, suffmul))
		if nums[i] == 0 {
			premul = 1
		}
		if nums[n-i-1] == 0 {
			suffmul = 1
		}
		println("Prefix and Suffix: ", premul, " ", suffmul)
	}
	return maxProd
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 0, -1, 2, 3, -5, -2}
	fmt.Printf("Res: %v", maxProduct(nums))
}

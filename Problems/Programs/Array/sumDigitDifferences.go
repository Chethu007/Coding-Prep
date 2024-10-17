package main

import "fmt"

func sumDigitDifferences(nums []int) int64 {
	freq := make([][]int, 10)
	for i := range freq {
		freq[i] = make([]int, 10)
	}
	//var sum int64
	//sum = 0
	for _, val := range nums {
		for i := 0; val != 0; i++ {
			rem := val % 10
			freq[i][rem]++
			val = val / 10
		}
	}
	fmt.Printf("Matrix: %v", freq)
	return 0
}

func main() {
	nums := []int{1234567, 1223243, 1234556}
	println("Res: ", sumDigitDifferences(nums))
}

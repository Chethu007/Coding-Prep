package main

import "fmt"

func generate(numRows int) [][]int {
	var ans [][]int
	for i := 1; i <= numRows; i++ {
		ans = append(ans, generateone(i))
	}
	return ans
}

func generateone(n int) []int {
	var res []int
	temp := 1
	res = append(res, temp)
	for i := 1; i < n; i++ {
		temp *= n - i
		temp = temp / i
		res = append(res, temp)
	}
	return res
}

func main() {
	fmt.Printf("Res: %v", generate(6))
}

package main

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	expectedMean := mean * (m + n)
	sum := 0
	for _, val := range rolls {
		sum += val
	}
	diff := expectedMean - sum
	if diff < 0 || diff/n == 0 || diff > 6*n {
		return []int{}
	}
	avg := diff / n
	res := make([]int, n)
	for i := 0; i < n-1; i++ {
		res[i] = avg
	}
	res[n-1] = diff - avg*(n-1)
	return res
}

func main() {
	fmt.Printf("Res: %v\n", missingRolls([]int{6, 6, 6, 6}, 1, 2))
}

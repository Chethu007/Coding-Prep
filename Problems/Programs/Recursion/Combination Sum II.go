package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)

	var solve func(i, sum int, temp []int, res *[][]int)
	solve = func(i, sum int, temp []int, res *[][]int) {
		if sum == 0 {
			*res = append(*res, append([]int{}, temp...))
			return
		}
		if i == n || sum < 0 {
			return
		}
		temp = append(temp, candidates[i])
		solve(i+1, sum-candidates[i], temp, res)
		temp = temp[:len(temp)-1]

		for i+1 < n && candidates[i+1] == candidates[i] {
			i++
		}
		solve(i+1, sum, temp, res)
	}

	var res [][]int
	solve(0, target, []int{}, &res)
	return res
}

func main() {
	fmt.Printf("Res: %v\n", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

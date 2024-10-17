package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rmin, rmax, lmin, lmax := 1001, -1, 1001, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rmin = min(rmin, i)
				rmax = max1(rmax, i)
				lmin = min(lmin, j)
				lmax = max1(lmax, j)
			}
		}
	}
	return (rmax - rmin + 1) * (lmax - lmin + 1)
}

func main() {
	grid := [][]int{
		{0, 0},
		{1, 0},
	}
	fmt.Println(minimumArea(grid))
}

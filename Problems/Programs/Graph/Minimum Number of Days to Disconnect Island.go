package main

import (
	"fmt"
	"time"
)

type item1 struct {
	row int
	col int
}

func minDays(grid [][]int) int {
	rLen, cLen := len(grid), len(grid[0])
	visited := make([][]int, rLen)
	for i := 0; i < rLen; i++ {
		visited[i] = make([]int, cLen)
	}
	// DFS Function
	var dfs func(row, col int, visited [][]int)
	dfs = func(row, col int, visited [][]int) {
		if row < 0 || col < 0 || row >= rLen || col >= cLen ||
			visited[row][col] != 0 || grid[row][col] != 1 {
			return
		}
		visited[row][col] = 1
		dfs(row+1, col, visited)
		dfs(row-1, col, visited)
		dfs(row, col+1, visited)
		dfs(row, col-1, visited)
	}

	var land []item1
	isLand := 0
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if grid[i][j] == 1 {
				land = append(land, item1{i, j})
			}
			if grid[i][j] == 1 && visited[i][j] == 0 {
				isLand++
				dfs(i, j, visited)
				if isLand > 1 {
					return 0
				}
			}
		}
	}
	if isLand == 0 {
		return 0
	}
	for _, val := range land {
		grid[val.row][val.col] = 0
		isLand1 := 0
		visited1 := make([][]int, rLen)
		for i := 0; i < rLen; i++ {
			visited1[i] = make([]int, cLen)
		}
		for i := 0; i < rLen; i++ {
			for j := 0; j < cLen; j++ {
				if grid[i][j] == 1 && visited1[i][j] == 0 {
					isLand1++
					dfs(i, j, visited1)
				}
				if isLand1 > 1 {
					return 1
				}
			}
		}
		if isLand1 == 0 {
			return 1
		}
		grid[val.row][val.col] = 1
	}
	return 2
}

func main() {
	grid := [][]int{{1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 0, 1, 1}, {1, 1, 0, 1, 1}}
	start := time.Now()
	fmt.Printf("Res: %v ", minDays(grid))
	elapsed := time.Since(start)
	fmt.Printf("Function 1 execution time: %s\n", elapsed)
}

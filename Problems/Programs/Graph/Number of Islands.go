package main

import (
	"fmt"
	"time"
)

type data struct {
	Row int
	Col int
}

// BFS
func numIslands(grid [][]byte) int {
	rLen, cLen := len(grid), len(grid[0])
	visited := make([][]int, rLen)
	for i := 0; i < rLen; i++ {
		visited[i] = make([]int, cLen)
	}
	isLand := 0

	var bfs func(row, col int, visited [][]int)
	bfs = func(row, col int, visited [][]int) {
		visited[row][col] = 1
		var q []data
		q = append(q, data{row, col})
		for len(q) > 0 {
			delrow := []int{-1, 1, 0, 0}
			delcol := []int{0, 0, -1, 1}
			cur := q[0]
			q = q[1:]
			row = cur.Row
			col = cur.Col
			//An island is surrounded by water and is formed by connecting adjacent
			// lands horizontally or vertically.
			// Option 1
			//if row-1 >= 0 && grid[row-1][col] == 1 && visited[row-1][col] == 0 {
			//	visited[row-1][col] = 1
			//	q = append(q, data{row - 1, col})
			//}
			//if row+1 < rLen && grid[row+1][col] == 1 && visited[row+1][col] == 0 {
			//	visited[row+1][col] = 1
			//	q = append(q, data{row + 1, col})
			//}
			//if col-1 >= 0 && grid[row][col-1] == 1 && visited[row][col-1] == 0 {
			//	visited[row][col-1] = 1
			//	q = append(q, data{row, col - 1})
			//}
			//if col+1 < cLen && grid[row][col+1] == 1 && visited[row][col+1] == 0 {
			//	visited[row][col+1] = 1
			//	q = append(q, data{row, col + 1})
			//}

			//Option 2
			for i := 0; i < 4; i++ {
				nrow := row + delrow[i]
				ncol := col + delcol[i]
				if nrow >= 0 && nrow < rLen && ncol >= 0 && ncol < cLen &&
					grid[nrow][ncol] == 1 && visited[nrow][ncol] == 0 {
					visited[nrow][ncol] = 1
					q = append(q, data{nrow, ncol})
				}
			}
			//Incase, island is surrounded by water and is formed by connecting adjacent
			// lands horizontally or vertically or diagonally.
			//for delrow := -1; delrow <= 1; delrow++ {
			//	for delcol := -1; delcol <= 1; delcol++ {
			//		nrow := row + delrow
			//		ncol := col + delcol
			//		if nrow >= 0 && nrow < rLen && ncol >= 0 && ncol < cLen &&
			//			grid[nrow][ncol] == 1 && visited[nrow][ncol] == 0 {
			//			visited[nrow][ncol] = 1
			//			q = append(q, data{nrow, ncol})
			//		}
			//	}
			//}
		}
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				isLand++
				bfs(i, j, visited)
			}
		}
	}
	return isLand
}

// BFS without visited matrix
func numIslandsNoVisited(grid [][]byte) int {
	rLen, cLen := len(grid), len(grid[0])
	isLand := 0
	var bfs func(row, col int)
	bfs = func(row, col int) {
		grid[row][col] = 0
		var q []data
		q = append(q, data{row, col})
		for len(q) > 0 {
			delrow := []int{-1, 1, 0, 0}
			delcol := []int{0, 0, -1, 1}
			cur := q[0]
			q = q[1:]
			row = cur.Row
			col = cur.Col
			for i := 0; i < 4; i++ {
				nrow := row + delrow[i]
				ncol := col + delcol[i]
				if nrow >= 0 && nrow < rLen && ncol >= 0 && ncol < cLen &&
					grid[nrow][ncol] == 1 {
					grid[nrow][ncol] = 0
					q = append(q, data{nrow, ncol})
				}
			}
		}
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if grid[i][j] == 1 {
				isLand++
				bfs(i, j)
			}
		}
	}
	return isLand
}

// DFS
func numIslands1(grid [][]byte) int {
	rLen, cLen := len(grid), len(grid[0])
	isLand := 0
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rLen || col < 0 || col >= cLen || grid[row][col] == 0 {
			return
		}
		grid[row][col] = 0
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if grid[i][j] == 1 {
				isLand++
				dfs(i, j)
			}
		}
	}
	return isLand
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	start := time.Now()
	fmt.Printf("Res: %v ", numIslands(grid))
	elapsed := time.Since(start)
	fmt.Printf("Function 1 execution time: %s\n", elapsed)
	//start = time.Now()
	//fmt.Printf("Res: %v ", numIslands1(grid))
	//elapsed = time.Since(start)
	//fmt.Printf("Function 2 execution time: %s\n", elapsed)
	start = time.Now()
	fmt.Printf("Res: %v ", numIslandsNoVisited(grid))
	elapsed = time.Since(start)
	fmt.Printf("Function 3 execution time: %s\n", elapsed)

}

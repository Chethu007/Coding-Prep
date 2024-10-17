package main

import (
	"fmt"
	"time"
)

type item struct {
	Row int
	Col int
	Min int
}

func orangesRotting(grid [][]int) int {
	rLen, cLen := len(grid), len(grid[0])
	count, count1 := 0, 0
	visited := make([][]int, rLen)
	for i := 0; i < rLen; i++ {
		visited[i] = make([]int, cLen)
	}
	var q []item
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if grid[i][j] == 2 {
				q = append(q, item{i, j, 0})
				visited[i][j] = 2
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}
	minTime := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		curRow := cur.Row
		curCol := cur.Col
		curMin := cur.Min
		if minTime < curMin {
			minTime = curMin
		}
		delRow := []int{1, -1, 0, 0}
		delCol := []int{0, 0, 1, -1}
		for i := 0; i < 4; i++ {
			nrow := curRow + delRow[i]
			ncol := curCol + delCol[i]
			if nrow >= 0 && ncol >= 0 && nrow < rLen && ncol < cLen &&
				grid[nrow][ncol] == 1 && visited[nrow][ncol] == 0 {
				visited[nrow][ncol] = 2
				q = append(q, item{nrow, ncol, curMin + 1})
				count1++
			}
		}
	}
	if count != count1 {
		return -1
	}
	return minTime
}

func main() {
	start := time.Now()
	fmt.Printf("Res: %v\n", orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Printf("Time take %v\n", time.Since(start))
	start = time.Now()
	fmt.Printf("Res: %v\n", orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Printf("Time take %v\n", time.Since(start))
	start = time.Now()
	fmt.Printf("Res: %v\n", orangesRotting([][]int{{0, 2}}))
	fmt.Printf("Time take %v\n", time.Since(start))
}

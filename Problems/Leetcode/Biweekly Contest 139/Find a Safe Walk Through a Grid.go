package main

import "fmt"

type data struct {
	Row       int
	Col       int
	curHealth int
}

func findSafeWalk(grid [][]int, health int) bool {
	rLen, cLen := len(grid), len(grid[0])
	visited := make([][]int, rLen)
	for i := 0; i < rLen; i++ {
		visited[i] = make([]int, cLen)
	}
	delrow := []int{-1, 1, 0, 0}
	delcol := []int{0, 0, -1, 1}
	initialHealth := health - grid[0][0]
	var q []data
	visited[0][0] = initialHealth
	q = append(q, data{0, 0, initialHealth})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		row := cur.Row
		col := cur.Col
		curHealth := cur.curHealth
		if row == rLen-1 && col == cLen-1 && curHealth > 0 {
			return true
		}
		for i := 0; i < 4; i++ {
			nrow := row + delrow[i]
			ncol := col + delcol[i]
			if nrow >= 0 && nrow < rLen && ncol >= 0 && ncol < cLen {
				nextHealth := curHealth - grid[nrow][ncol]
				if nextHealth > 0 && nextHealth > visited[nrow][ncol] {
					visited[nrow][ncol] = nextHealth
					q = append(q, data{nrow, ncol, nextHealth})
				}
			}
		}
	}
	return false
}

func main() {
	grid := [][]int{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 1},
		{0, 0, 1, 0, 1, 0},
	}
	fmt.Printf("Res: %v", findSafeWalk(grid, 4))
}

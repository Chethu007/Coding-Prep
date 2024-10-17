package main

import (
	"fmt"
	"time"
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rLen, cLen := len(image), len(image[0])
	startingPixel := image[sr][sc]
	if startingPixel == color {
		return image
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rLen || col < 0 || col >= cLen || image[row][col] != startingPixel {
			return
		}
		image[row][col] = color
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}
	dfs(sr, sc)
	return image
}

func main() {
	//grid := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	//start := time.Now()
	//fmt.Printf("Res: %v\n", floodFill(grid, 1, 1, 2))
	//elapsed := time.Since(start)
	//fmt.Printf("Function completed in : %s seconds\n", elapsed)
	grid := [][]int{{0, 0, 0}, {0, 0, 0}}
	start := time.Now()
	fmt.Printf("Res: %v\n", floodFill(grid, 0, 0, 0))
	elapsed := time.Since(start)
	fmt.Printf("Function completed in : %s seconds\n", elapsed)
}

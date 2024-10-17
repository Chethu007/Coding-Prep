package main

import (
	"fmt"
	"strings"
)

//func numMagicSquaresInside(grid [][]int) int {
//	rLen, cLen := len(grid), len(grid[0])
//	if rLen < 3 || cLen < 3 {
//		return 0
//	}
//	var solve func(row, col int) bool
//	solve = func(row, col int) bool {
//		m := make(map[int]int)
//		for i := row; i < row+3; i++ {
//			for j := col; j < col+3; j++ {
//				if grid[i][j] < 1 || grid[i][j] > 9 {
//					return false
//				} else {
//					if m[grid[i][j]] == 0 {
//						m[grid[i][j]] = 1
//					} else {
//						return false
//					}
//				}
//			}
//		}
//		rSum1 := grid[row][col] + grid[row][col+1] + grid[row][col+2]
//		rSum2 := grid[row+1][col] + grid[row+1][col+1] + grid[row+1][col+2]
//		rSum3 := grid[row+2][col] + grid[row+2][col+1] + grid[row+2][col+2]
//		cSum1 := grid[row][col] + grid[row+1][col] + grid[row+2][col]
//		cSum2 := grid[row][col+1] + grid[row+1][col+1] + grid[row+2][col+1]
//		cSum3 := grid[row][col+2] + grid[row+1][col+2] + grid[row+2][col+2]
//		dSum1 := grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
//		dSum2 := grid[row+2][col] + grid[row+1][col+1] + grid[row][col+2]
//		if rSum1 == rSum2 && rSum2 == rSum3 && rSum3 == cSum1 && cSum1 == cSum2 && cSum2 == cSum3 && cSum3 == dSum1 && dSum1 == dSum2 {
//			return true
//		}
//		return false
//	}
//	count := 0
//	for i := 0; i+2 < rLen; i++ {
//		for j := 0; j+2 < cLen; j++ {
//			if solve(i, j) {
//				count++
//			}
//		}
//	}
//	return count
//}

func numMagicSquaresInside(grid [][]int) int {
	rLen, cLen := len(grid), len(grid[0])
	if rLen < 3 || cLen < 3 {
		return 0
	}
	var solve func(row, col int) bool
	solve = func(row, col int) bool {
		patten1 := "4381672943816729"
		patten2 := "4927618349276183"
		if grid[row][col] != 5 {
			return false
		}
		m := map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
		str := ""
		str += m[grid[row-1][col-1]] + m[grid[row-1][col]] + m[grid[row-1][col+1]] + m[grid[row][col+1]]
		str += m[grid[row+1][col+1]] + m[grid[row+1][col]] + m[grid[row+1][col-1]] + m[grid[row][col-1]]
		if strings.Contains(patten1, str) || strings.Contains(patten2, str) {
			return true
		}
		return false
	}
	count := 0
	for i := 1; i < rLen-1; i++ {
		for j := 1; j < cLen-1; j++ {
			if solve(i, j) {
				count++
			}
		}
	}
	return count
}

func main() {
	grid := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
	fmt.Printf("Res: %v\n", numMagicSquaresInside(grid))
}

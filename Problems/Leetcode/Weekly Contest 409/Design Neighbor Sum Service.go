package main

import "fmt"

type neighborSum struct {
	Grid [][]int
}

func Constructor(grid [][]int) neighborSum {
	rLen, cLen := len(grid), len(grid[0])
	matrix := make([][]int, rLen)
	for i := 0; i < rLen; i++ {
		matrix[i] = make([]int, cLen)
	}
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			matrix[i][j] = grid[i][j]
		}
	}
	return neighborSum{Grid: matrix}
}

func (this *neighborSum) AdjacentSum(value int) int {
	matrix := this.Grid
	rLen, cLen := len(matrix), len(matrix[0])
	sum := 0
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if matrix[i][j] == value {
				if j+1 < cLen {
					sum += matrix[i][j+1]
				}
				if j-1 >= 0 {
					sum += matrix[i][j-1]
				}
				if i-1 >= 0 {
					sum += matrix[i-1][j]
				}
				if i+1 < rLen {
					sum += matrix[i+1][j]
				}
				break
			}
		}
	}
	return sum
}

func (this *neighborSum) DiagonalSum(value int) int {
	matrix := this.Grid
	rLen, cLen := len(matrix), len(matrix[0])
	sum := 0
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if matrix[i][j] == value {
				if j+1 < cLen && i+1 < rLen {
					println("Sum", sum, "matrix", matrix[i+1][j+1])
					sum += matrix[i+1][j+1]
				}
				if j-1 >= 0 && i+1 < rLen {
					println("Sum", sum, "matrix", matrix[i+1][j-1])
					sum += matrix[i+1][j-1]
				}
				if i-1 >= 0 && j+1 < cLen {
					println("Sum", sum, "matrix", matrix[i-1][j+1])
					sum += matrix[i-1][j+1]
				}
				if i-1 >= 0 && j-1 >= 0 {
					println("Sum", sum, "matrix", matrix[i+1][j-1])
					sum += matrix[i-1][j-1]
				}
				break
			}
		}
	}
	return sum
}

/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */

func main() {
	grid := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	obj := Constructor(grid)
	fmt.Printf("Matrix: %v", obj)
	param_1 := obj.AdjacentSum(4)
	fmt.Println(param_1)
	param_2 := obj.DiagonalSum(3)
	fmt.Println(param_2)
}

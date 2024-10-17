package main

import "fmt"

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
func main() {
	source := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 5},
		{7, 8, 9, 6},
		{7, 8, 9, 6},
	}

	fmt.Println("Source matrix:")
	printMatrix(source)
	//rotate(source)
	//fmt.Println("Destination matrix:")
	//printMatrix(source)
	transpose(source)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = matrix[n-1-j][i]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = res[i][j]
		}
	}
}

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if i > j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	fmt.Println("Transpose matrix:")
	printMatrix(matrix)
}

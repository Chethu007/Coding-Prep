package main

import "fmt"

//	func solve(board [][]byte) {
//		rLen, cLen := len(board), len(board[0])
//
//		var replace func(r int)
//		replace = func(r int) {
//			var char byte
//			if r == 1 {
//				char = 'O'
//			} else {
//				char = 'X'
//			}
//			for i := 0; i < rLen; i++ {
//				for j := 0; j < cLen; j++ {
//					if board[i][j] == 'R' {
//						board[i][j] = char
//					}
//				}
//			}
//		}
//
//		var dfs func(row, col int, res *int) int
//		dfs = func(row, col int, res *int) int {
//			if row < 0 || col < 0 || row >= rLen || col >= cLen || board[row][col] != 'O' {
//				return *res
//			}
//			if row == 0 || col == 0 || row == rLen-1 || col == cLen-1 {
//				*res = 1
//			}
//			board[row][col] = 'R'
//			dfs(row+1, col, res)
//			dfs(row-1, col, res)
//			dfs(row, col-1, res)
//			dfs(row, col+1, res)
//			return *res
//		}
//		var res int
//		for i := 0; i < rLen; i++ {
//			for j := 0; j < cLen; j++ {
//				if board[i][j] == 'O' {
//					dfs(i, j, &res)
//					replace(res)
//				}
//				res = 0
//			}
//		}
//	}
func solve(board [][]byte) {
	rLen, cLen := len(board), len(board[0])
	visited := make([][]int, rLen)
	for i := 0; i < rLen; i++ {
		visited[i] = make([]int, cLen)
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= rLen || col >= cLen || board[row][col] != 'O' || visited[row][col] == 1 {
			return
		}
		visited[row][col] = 1
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if (i == 0 || j == 0 || i == rLen-1 || j == cLen-1) && visited[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if board[i][j] == 'O' && visited[i][j] == 0 {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Printf("Res : %c\n", board)
}

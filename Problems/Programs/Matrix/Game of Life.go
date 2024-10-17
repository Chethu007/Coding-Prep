package main

import "fmt"

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	var checkLiveCell func(i, j int) int
	checkLiveCell = func(i, j int) int {
		count := 0
		// down
		if i+1 != m {
			count += board[i+1][j]
			if j+1 != n {
				count += board[i+1][j+1]
			}
			if j-1 != -1 {
				count += board[i+1][j-1]
			}
		}
		// up
		if i-1 != -1 {
			count += board[i-1][j]
			if j+1 != n {
				count += board[i-1][j+1]
			}
			if j-1 != -1 {
				count += board[i-1][j-1]
			}
		}
		// left
		if j-1 != -1 {
			count += board[i][j-1]
		}
		//right
		if j+1 != n {
			count += board[i][j+1]
		}
		return count
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveCell := checkLiveCell(i, j)
			res[i][j] = board[i][j]
			if board[i][j] == 0 {
				if liveCell == 3 {
					res[i][j] = 1
				}
			} else {
				if liveCell > 3 || liveCell < 2 {
					res[i][j] = 0
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = res[i][j]
		}
	}
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	//board := [][]int{
	//	{1, 1},
	//	{1, 0},
	//}
	gameOfLife(board)
	fmt.Println(board)
}

package main

func minFlips(grid [][]int) int {
	rLen, cLen := len(grid), len(grid[0])
	rCount, cCount := 0, 0
	for i := 0; i < rLen; i++ {
		rCount += isRowPalindrome(grid[i], cLen)
	}
	for i := 0; i < cLen; i++ {
		cCount += isColPalindrome(i, rLen, grid)
	}
	if rCount < cCount {
		return rCount
	}
	return cCount
}

func isRowPalindrome(row []int, n int) int {
	count := 0
	for i := 0; i < n/2; i++ {
		if row[i] != row[n-i-1] {
			count++
		}
	}
	return count
}

func isColPalindrome(j, n int, grid [][]int) int {
	count := 0
	for i := 0; i < n/2; i++ {
		if grid[i][j] != grid[n-i-1][j] {
			count++
		}
	}
	return count
}

func main() {
	println(minFlips([][]int{{1}, {0}}))
}

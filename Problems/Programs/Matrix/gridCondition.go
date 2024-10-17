package main

func satisfiesConditions(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	if m == 1 && n > 1 {
		for i := 0; i < n-1; i++ {
			if grid[0][i] == grid[0][i+1] {
				return false
			}
		}
	} else if m > 1 && n == 1 {
		for i := 0; i < m-1; i++ {
			if grid[i][0] != grid[i+1][0] {
				return false
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j+1 < n {
				if grid[i][j] == grid[i][j+1] {
					return false
				}
			}
			if i+1 < m {
				if grid[i][j] != grid[i+1][j] {
					return false
				}
			}

		}
	}
	return true
}

func main() {
	grid := [][]int{{1, 0, 2}, {1, 0, 2}}
	println("Res: ", satisfiesConditions(grid))
}

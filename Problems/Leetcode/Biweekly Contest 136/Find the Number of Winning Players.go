package main

import (
	"sort"
)

func winningPlayerCount(n int, pick [][]int) int {
	sort.Slice(pick, func(i, j int) bool {
		if pick[i][0] == pick[j][0] {
			return pick[i][1] < pick[j][1]
		}
		return pick[i][0] < pick[j][0]
	})
	count, i, length := 0, 0, len(pick)
	if pick[0][0] == 0 {
		count++
	}
	for i < length && pick[i][0] == 0 {
		i++
	}
	for j := 1; j < n; j++ {
		//println("j Loop", j)
		m := make(map[int]int)
		for i < length && pick[i][0] == j {
			m[pick[i][1]]++
			//println("pick[i][0]", pick[i][0], "pick[i][1]", pick[i][1], "i", i)
			i++
		}
		for _, val := range m {
			if val >= j+1 {
				count++
				break
			}
		}
		if i == length {
			return count
		}
	}
	return count
}

func main() {
	println(winningPlayerCount(4, [][]int{{0, 2}, {0, 7}, {1, 5}, {0, 5}, {1, 2}, {0, 6}, {1, 2}, {1, 5}, {1, 4}, {0, 7}}))
}

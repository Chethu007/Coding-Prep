package main

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})
	//fmt.Printf("Meeting %v\n", meetings)
	var res [][]int
	res = append(res, meetings[0])
	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] <= res[len(res)-1][1] && res[len(res)-1][1] < meetings[i][1] {
			res[len(res)-1][1] = meetings[i][1]
		} else if res[len(res)-1][1] < meetings[i][0] {
			res = append(res, meetings[i])
		}
	}
	//fmt.Printf("Meeting %v\n", res)
	n := len(res)
	result := res[0][0] - 1
	for i := 0; i < n-1; i++ {
		if res[i][1]+1 < res[i+1][0] {
			result += res[i+1][0] - res[i][1] - 1
		}
	}
	result += days - res[n-1][1]
	return result
}

func main() {
	data := [][]int{
		{1, 3},
		{2, 4},
	}
	fmt.Printf("Res: %v", countDays(5, data))
}

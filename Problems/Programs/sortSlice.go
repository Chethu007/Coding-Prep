package main

import (
	"fmt"
	"sort"
)

func main() {
	x := [][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}, {1, 5}, {2, 3}}

	// Define a custom sorting function
	comp := func(i, j int) bool {
		// Compare first elements
		if x[i][0] == x[j][0] {
			// If first elements are equal, compare second elements
			return x[i][1] < x[j][1]
		}
		return x[i][0] < x[j][0]
	}

	sort.Slice(x, comp)

	fmt.Println("Sorted:", x)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	//var res [][]int
	n := len(intervals)
	for i := 0; i < n; i++ {
		if intervals[i][0] >= newInterval[0] {
			intervals = append(intervals, []int{0, 0})
			copy(intervals[i+1:], intervals[i:])
			intervals[i] = newInterval
			break
		}
	}
	return intervals
}

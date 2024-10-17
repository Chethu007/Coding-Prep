package main

import "fmt"

func copy1(target, source [][]int) {
	println(len(target), " ", len(source))
	minLen := len(target)
	if len(source) < minLen {
		minLen = len(source)
	}

	for i := 0; i < minLen; i++ {
		target[i] = source[i]
	}
}

func main() {
	// Original intervals
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	fmt.Println("Original intervals:", intervals)

	intervals = append(intervals, []int{0, 0})
	var newInterval [][]int
	for _, list := range intervals {
		newInterval = append(newInterval, list)
	}

	// Copy modified intervals from source to target
	copy1(intervals[2:], newInterval[1:])

	// Print the result
	fmt.Println("Modified intervals:", intervals)
}

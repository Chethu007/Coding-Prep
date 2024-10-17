package main

import "fmt"

func numberOfAlternatingGroups(colors []int) int {
	count, n := 0, len(colors)
	for i := 1; i < n-1; i++ {
		if colors[i-1] != colors[i] && colors[i] != colors[i+1] {
			count++
		}
	}
	if colors[0] != colors[1] && colors[0] != colors[n-1] {
		count++
	}
	if colors[0] != colors[n-1] && colors[n-1] != colors[n-2] {
		count++
	}
	return count
}

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1}))
}

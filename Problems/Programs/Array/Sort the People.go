package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	var res []string
	m := make(map[int]string)
	for i := range names {
		m[heights[i]] = names[i]
	}
	sort.Ints(heights)
	for i := range heights {
		println(heights[i], m[heights[i]])
		res = append(res, m[heights[i]])
	}
	return res
}

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	println(sortPeople(names, heights))
}

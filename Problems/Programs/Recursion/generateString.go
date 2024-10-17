package main

import "fmt"

func main() {
	var res []string
	var solve func(index, n int, s string, res *[]string)
	solve = func(index, n int, s string, res *[]string) {
		if index == n {
			*res = append(*res, s)
			return
		}
		solve(index+1, n, s+string('0'), res)
		if len(s) == 0 || s[len(s)-1] != '1' {
			solve(index+1, n, s+string('1'), res)
		}
	}
	n := 4
	solve(0, n, "", &res)
	fmt.Printf("Res: %v", res)
}

package main

import "fmt"

func validStrings(n int) []string {
	var res []string
	var solve func(i int, temp string, res *[]string)
	solve = func(i int, temp string, res *[]string) {
		if i == n {
			*res = append(*res, temp)
			return
		}
		if len(temp) > 0 && temp[len(temp)-1] == '0' {
			solve(i+1, temp+"1", res)
		} else {
			solve(i+1, temp+"0", res)
			solve(i+1, temp+"1", res)
		}
	}
	solve(0, "", &res)
	return res
}

func main() {
	fmt.Println(validStrings(2))
}

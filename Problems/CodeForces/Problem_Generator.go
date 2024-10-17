package main

import (
	"fmt"
)

func solve(str string, m int) int {
	var res int
	arr := [7]int{}
	for i := range str {
		arr[str[i]-'A']++
	}
	for i := range arr {
		if arr[i] < m {
			res += m - arr[i]
		}
	}
	return res
}

func main() {
	var t int
	fmt.Scanln(&t)

	for t > 0 {
		var n, m int
		fmt.Scanln(&n, &m)
		var str string
		fmt.Scanln(&str)
		result := solve(str, m)

		fmt.Println(result)
		t--
	}
}

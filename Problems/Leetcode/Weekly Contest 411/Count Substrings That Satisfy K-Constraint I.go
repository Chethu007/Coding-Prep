package main

import "fmt"

func countKConstraintSubstrings(s string, k int) int {
	n, count := len(s), 0

	var solve func(str string) bool
	solve = func(str string) bool {
		zero, one := 0, 0
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				zero++
			} else {
				one++
			}
		}
		if zero > k && one > k {
			return false
		}
		return true
	}

	for i := 0; i < n; i++ {
		count++
		str := string(s[i])
		for j := i + 1; j < n; j++ {
			if j-i == 1 || j-i == 2 {
				str += string(s[j])
				count++
				continue
			}
			str += string(s[j])
			if solve(str) {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Printf("Res: %v\n", countKConstraintSubstrings("1010101", 2))
}

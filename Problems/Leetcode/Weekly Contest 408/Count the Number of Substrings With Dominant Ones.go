package main

import "fmt"

func numberOfSubstrings(s string) int {
	count := 0
	n := len(s)
	for i := 0; i < n; i++ {
		one, zero := 0, 0
		if s[i] == '1' {
			one++
			count++
		} else {
			zero++
		}
		for j := i + 1; j < n; j++ {
			if s[j] == '1' {
				one++
			} else {
				zero++
			}
			if one >= zero*zero {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(numberOfSubstrings("101101"))
}

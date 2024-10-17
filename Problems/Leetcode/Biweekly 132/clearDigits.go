package main

import "fmt"

func clearDigits(s string) string {
	res := []byte{}
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			res = append(res, s[i])
		} else {
			if len(res) != 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return string(res)
}

func main() {
	s := "cab34a"
	fmt.Println(clearDigits(s))
}

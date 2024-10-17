package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	str := strings.Split(s, " ")
	var temp []string
	for _, val := range str {
		if strings.TrimSpace(val) != "" {
			temp = append(temp, val)
		}
	}
	res := ""
	for i := len(temp) - 1; i >= 0; i-- {
		res += temp[i] + " "
	}
	return res[:len(res)-1]
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}

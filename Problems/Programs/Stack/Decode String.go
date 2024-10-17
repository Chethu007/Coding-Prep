package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var st []byte
	for i := range s {
		if s[i] != ']' {
			st = append(st, s[i])
		} else {
			substr, k := "", ""
			top := len(st) - 1
			for st[top] != '[' {
				substr = string(st[top]) + substr
				top--
			}
			st = st[:top]
			top = len(st) - 1
			for top >= 0 && st[top] >= '0' && st[top] <= '9' {
				k = string(st[top]) + k
				top--
			}
			st = st[:top+1]
			kval := 1
			if k != "" {
				kval, _ = strconv.Atoi(k)
			}
			substr = strings.Repeat(substr, kval)
			println("Substring", substr, "Kval", kval, "k", k)
			st = append(st, []byte(substr)...)
		}
	}
	return string(st)
}

func main() {
	println(decodeString("3[a]2[bc]"))
}

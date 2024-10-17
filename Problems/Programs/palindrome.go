package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	println(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isalphanumeric(rune(s[i])) {
			i++
			continue
		}
		if !isalphanumeric(rune(s[j])) {
			j--
			continue
		}
		println(s[i], " ", s[j])
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isalphanumeric(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}

func main() {
	println(isPalindrome("A man, a plan, a canal: Panama"))
	println(isPalindrome("0P"))
	println(isPalindrome("09P90"))
}

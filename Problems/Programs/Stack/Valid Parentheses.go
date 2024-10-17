package main

import "fmt"

func isValid(s string) bool {
	i, n := 0, len(s)
	var st []byte
	//if s[0] == '}' || s[0] == ')' || s[0] == ']' {
	//	return false
	//}
	for i < n {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			st = append(st, s[i])
		} else {
			if len(st) != 0 {
				if (s[i] == ')' && st[len(st)-1] == '(') || (s[i] == '}' && st[len(st)-1] == '{') || (s[i] == ']' && st[len(st)-1] == '[') {
					st = st[:len(st)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
		//fmt.Printf("Stack: %v\n", string(st))
		i++
	}
	if len(st) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("[]"))
}

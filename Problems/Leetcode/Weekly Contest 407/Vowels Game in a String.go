package main

func doesAliceWin(s string) bool {
	count := 0
	for i := range s {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			count++
		}
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	println(doesAliceWin("aaaaaq"))
}

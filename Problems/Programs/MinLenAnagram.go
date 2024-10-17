package main

func minAnagramLength(s string) int {
	n := len(s)
	res := 0
	m := make([]int, 26)
	for i := range s {
		m[s[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if m[i] == 0 {
			continue
		}
		if n%m[i] == 0 {
			println(m[i])
			res++
		} else {
			return n
		}
	}
	return res
}

func main() {
	println(minAnagramLength("jjj"))
}

package main

func minimumLength(s string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	println(len(m), len(s))
	res := 0
	for _, val := range m {
		res += (val - 1) / 2
	}
	return len(s) - (2 * res)
}

func main() {
	println(minimumLength("aaaaaaabbbbbbbcccccccdef"))
}

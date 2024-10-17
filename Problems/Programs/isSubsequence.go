package main

func isSubsequence(s string, t string) bool {
	substr := []rune(s)
	targetstr := []rune(t)
	if len(substr) == 0 {
		return true
	}
	i, j, count := 0, 0, 0
	for i < len(targetstr) {
		if targetstr[i] == substr[j] {
			count++
			i++
			j++
		}
		i++
	}
	if count == len(substr) {
		return true
	}
	return false
}

func main() {
	println(isSubsequence("abx", "ksjfkoa"))
}

package main

func maxOperations(s string) int {
	n, count, res := len(s), 0, 0
	count1 := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			count++
			count1++
		} else if count > 0 {
			res += count1
			count = 0
		}
		println(i, count, res)
	}
	return res
}

func main() {
	println(maxOperations("00111"))
}

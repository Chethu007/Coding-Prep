package main

func stringHash(s string, k int) string {
	i, n, res := 0, len(s), ""
	for i < n {
		sum := 0
		for j := 0; j < k; j++ {
			sum += int(s[j+i] - 'a')
		}
		sum = sum % 26
		res += string('a' + rune(sum))
		i += k
	}
	return res
}

func main() {
	println("Res: ", stringHash("mxz", 3))
}

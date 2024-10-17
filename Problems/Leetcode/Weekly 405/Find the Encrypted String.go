package main

import "fmt"

func getEncryptedString(s string, k int) string {
	var res string
	n := len(s)
	for i := 0; i < n; i++ {
		res += string(s[(i+k)%n])
		println(string(s[i]), res)
	}
	return res
}

func main() {
	fmt.Println(getEncryptedString("dart", 3))
}

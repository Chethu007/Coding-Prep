package main

import "fmt"

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}
	m := make(map[int]string)
	l, k := 0, 0
	for l < n {
		if k == 0 {
			for i := 0; i < numRows && l < n; i++ {
				m[i] += string(s[l])
				l++
				k++
			}
		}
		if k == numRows {
			k -= 2
			for j := k; j > 0 && l < n; j-- {
				m[j] += string(s[l])
				l++
				k--
			}
		}

	}
	res := ""
	for i := 0; i < numRows; i++ {
		res += m[i]
		println(i, "->", m[i])
	}
	return res
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

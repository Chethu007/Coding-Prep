package main

import "fmt"

func minLength(s string) int {
	n := len(s)
	i := 0
	var temp string
	for {
		flag := false
		temp = s
		println("n = ", n, " temp = ", temp)
		for i = 0; i < n-1; i++ {
			if (temp[i] == 'A' && temp[i+1] == 'B') || (temp[i] == 'C' && temp[i+1] == 'D') {
				flag = true
				break
			}
		}
		println("i = ", i)
		if flag == false {
			return n
		}
		if i == 0 {
			s = temp[2:]
		} else if i == n-2 {
			s = temp[:n-2]
		} else {
			s = temp[:i] + temp[i+2:]
		}
		println("s = ", s)
		n -= 2
		if n == 0 {
			return 0
		}
		i = 0
	}
	return 0
}

func main() {
	fmt.Printf("Res: %v\n", minLength("CCCCDDDD"))
}

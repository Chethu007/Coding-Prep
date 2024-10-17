package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		if n >= m {
			fmt.Println(n - m)
		} else {
			x := m % n
			if x > n-x {
				x = n - x
			}
			fmt.Println(x)
		}
	}
}

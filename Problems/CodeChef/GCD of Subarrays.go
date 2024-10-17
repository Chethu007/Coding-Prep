package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var count int
	fmt.Fscanln(reader, &count)
	for j := 0; j < count; j++ {
		var n, k uint64
		fmt.Fscanln(reader, &n, &k)
		// Calculating the sum of the first n natural numbers
		total := n * (n + 1) / 2
		if total > k {
			fmt.Println(-1)
		} else {
			for i := uint64(0); i < n-1; i++ {
				fmt.Print(1, " ")
			}
			fmt.Println(k - total + 1) // using Println to ensure newline is added after the last number
		}
	}
}

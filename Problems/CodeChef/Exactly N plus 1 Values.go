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
		var n int
		fmt.Fscanln(reader, &n)
		fmt.Print(1, " ")
		for i := 0; i < n; i++ {
			fmt.Print(1<<i, " ")
		}
		fmt.Println()
	}
}

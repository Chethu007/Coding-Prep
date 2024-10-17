package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n, c int
		fmt.Fscan(reader, &n, &c)
		arr := make([]int, n)
		total := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &arr[j])
			total += arr[j]
		}
		sort.Ints(arr)
		del, x := 0, 0
		for j := 0; j < n; j++ {
			x += arr[j] * (total - arr[j])
			println("i: ", i, "x: ", x, "c: ", c)
			if x <= c {
				del++
				c -= x
			} else {
				break
			}
			total -= arr[j]
		}
		fmt.Println()
		if del == n {
			del--
		}
		fmt.Println("Res:", n-del)
	}
}

package main

import (
	"fmt"
	"math"
)

func Fibonaccharsis(i, n, k int64) bool {
	var j int64
	valid := true
	first := n
	second := i
	for j = 2; j < k; j++ {
		temp := second
		second = first - second
		first = temp
		if first < second || second < 0 || first < 0 {
			valid = false
			break
		}

	}
	return valid
}

func main() {
	var t, n, k, i int64
	fmt.Scan(&t)
	// Loop through each test case
	for t > 0 {
		// Read n
		var res int64 = 0
		fmt.Scan(&n, &k)
		if (k >= 7 && k >= n) || k > 30 {
			fmt.Println(0)
		} else {
			for i = int64(math.Ceil(float64(n) / 2)); i <= n; i++ {
				if Fibonaccharsis(i, n, k) {
					res++
				}
			}
			fmt.Println(res)
		}
		t--
	}
}

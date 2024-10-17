package main

import "fmt"

func valueAfterKSeconds(n int, k int) int {
	const MOD = 1e9 + 7
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	for i := 0; i < k; i++ {
		for j := 1; j < n; j++ {
			arr[j] += arr[j-1]
			arr[j] %= MOD
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(valueAfterKSeconds(5, 3))
}

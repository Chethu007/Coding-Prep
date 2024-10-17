package main

import (
	"fmt"
)

func nonSpecialCount(l int, r int) int {

	count := 0
	for i := 2; i*i <= r; i++ {
		if checkPrime(i) {
			sq := i * i
			if sq >= l && sq <= r {
				count++
			}
		}
	}
	//println(count)
	return r - l + 1 - count
}

func checkPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(nonSpecialCount(5, 7))
}

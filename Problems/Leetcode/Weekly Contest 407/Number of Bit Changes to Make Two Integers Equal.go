package main

func minChanges(n int, k int) int {
	if k > n {
		return -1
	}
	if n == k {
		return 0
	}
	count := 0
	for n > 0 || k > 0 {
		x := n & 1
		y := k & 1
		if y > x {
			return -1
		}
		if x > y {
			count++
		}
		n >>= 1
		k >>= 1
	}
	return count
}

func main() {
	println(minChanges(13, 14))
}

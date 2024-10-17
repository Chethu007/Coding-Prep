package main

import (
	"fmt"
	"math"
	"sort"
)

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func findClosestElements(arr []int, k int, x int) []int {
	n, minVal := len(arr), math.MaxInt32
	temp := make([]int, n)
	var minIndex int
	for i := 0; i < n; i++ {
		temp[i] = Min(absDiff(x, arr[i]), absDiff(arr[i], x))
		if minVal >= temp[i] {
			minVal = temp[i]
			minIndex = i
		}
	}
	res := make([]int, k)
	res[0] = arr[minIndex]
	println("MinIndex: ", minIndex, " and value: ", res[0])
	a, b := minIndex, minIndex
	j := 1
	flag := false
	if a > 0 && a < n-1 {
		flag = true
	}
	for j < k {
		if flag && a > 0 {
			a--
			res[j] = arr[a]
		} else if !flag && b < n {
			b++
			res[j] = arr[b]
		} else if a == 0 {
			b++
			res[j] = arr[b]
		} else {
			a--
			res[j] = arr[a]
		}
		j++
	}
	sort.Ints(res)
	return res
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func main() {
	n := []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}
	fmt.Printf("Res: %v", findClosestElements(n, 9, 4))
}

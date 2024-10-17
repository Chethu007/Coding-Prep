package main

import (
	"fmt"
)

func lostOperation() {
	var n, m int
	// Read n
	fmt.Scan(&n)
	arr := make([]int, n)
	newArr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&newArr[i])
	}
	fmt.Scan(&m)
	modArr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&modArr[i])
	}
	mp := make(map[int]int)
	for i := 0; i < m; i++ {
		mp[modArr[i]]++
	}
	flag := false
	for i := 0; i < n; i++ {
		if newArr[i] == modArr[m-1] {
			flag = true
		}
		if arr[i] != newArr[i] {
			if mp[newArr[i]] > 0 {
				mp[newArr[i]]--
				if mp[newArr[i]] == 0 {
					delete(mp, newArr[i])
				}
			} else {
				println("NO")
				return
			}
		}
	}
	fmt.Printf("%v", newArr)
	if !flag && len(mp) > 0 {
		println("NO")
		return
	}
	println("YES")
	return
}

func main() {
	var t int
	fmt.Scan(&t)
	// Loop through each test case
	for t > 0 {
		lostOperation()
		t--
	}
}

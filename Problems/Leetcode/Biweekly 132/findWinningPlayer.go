package main

import "fmt"

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	//pre := make([]int, n)
	//suf := make([]int, n)
	total := make([]int, n)
	//pre[n-1], suf[0] = 0, 0
	for i := 0; i < n; i++ {
		count1, count2 := 0, 0
		for j := i; j < n-1; j++ {
			if skills[i] > skills[j+1] {
				count1++
			} else {
				break
			}
		}
		for j := i; j >= 0; j-- {
			if j == 0 {
				continue
			}
			if skills[i] > skills[j-1] {
				count2++
				if i != n-1 {
					break
				}
			} else {
				break
			}
		}
		total[i] = count1 + count2
	}
	//fmt.Printf("Total : %v\n", total)
	maxVal, index := skills[0], 0
	for i := range skills {
		if maxVal < skills[i] {
			maxVal = skills[i]
			index = i
		}
	}
	//println("Index: ", index, " Value : ", maxVal)
	for i := range skills {
		if total[i] >= k || total[i] == n-1 {
			return i
		}
	}
	return index
}

func main() {
	nums := []int{8, 9, 7, 19, 11}
	k := 3
	fmt.Println(findWinningPlayer(nums, k))
}

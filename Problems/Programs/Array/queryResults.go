package main

import (
	"fmt"
)

func queryResults(limit int, queries [][]int) []int {
	var res []int
	mBall := make(map[int][3]int)
	mColor := make(map[int]int)
	count := 0
	for i := 0; i < len(queries); i++ {
		flag := false
		if mBall[queries[i][0]][0] == 0 {
			if mColor[queries[i][1]] == 0 {
				count++
				mColor[queries[i][1]]++
			}
		} else if mBall[queries[i][0]][0] == 1 {
			if mColor[queries[i][1]] == 0 {
				mColor[queries[i][1]]++
				mColor[mBall[queries[i][0]][1]]--
			} else {
				for _, arr := range mBall {
					//mBall[queries[i][0]][1] != queries[i][1] || mBall[queries[i][0]][1] != queries[i][0]
					if arr[1] == queries[i][1] && arr[2] == queries[i][0] {
						continue
					} else if arr[1] == mColor[queries[i][1]] {
						flag = true
					}
				}
			}
		}
		if flag == true {
			count--
		}
		mBall[queries[i][0]] = [3]int{1, queries[i][1], queries[i][1]}
		res = append(res, count)
	}
	return res
}

// [[1,4],[2,5],[1,3],[3,4]]
// Output: [1,2,2,3]
func main() {
	data := [][]int{
		{0, 1},
		{0, 4},
		{1, 2},
		{1, 5},
		{1, 4},
	}
	fmt.Printf("Res: %v", queryResults(1, data))
}

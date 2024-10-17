package main

import "fmt"

// O(m*n)
func maxPoints(points [][]int) int64 {
	rLen, cLen := len(points), len(points[0])
	curRow := make([]int64, cLen)
	nextRow := make([]int64, cLen)
	for i := 0; i < cLen; i++ {
		curRow[i] = int64(points[0][i])
	}
	for i := 1; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			nextRow[j] = int64(points[i][j])
		}
		left := make([]int64, cLen)
		right := make([]int64, cLen)
		left[0] = curRow[0]
		right[cLen-1] = curRow[cLen-1]
		for j := 1; j < cLen; j++ {
			left[j] = max(curRow[j], left[j-1]-1)
		}
		for j := cLen - 2; j >= 0; j-- {
			right[j] = max(curRow[j], right[j+1]-1)
		}
		for j := 0; j < cLen; j++ {
			nextRow[j] += max(left[j], right[j])
		}
		for j := 0; j < cLen; j++ {
			curRow[j] = nextRow[j]
		}
	}
	res := curRow[0]
	for i := 1; i < cLen; i++ {
		if res < curRow[i] {
			res = curRow[i]
		}
	}
	return res
}

// O(m*n*n)
//func maxPoints(points [][]int) int64 {
//	rLen, cLen := len(points), len(points[0])
//	curRow := make([]int64, cLen)
//	for i := 0; i < cLen; i++ {
//		curRow[i] = int64(points[0][i])
//	}
//	for i := 1; i < rLen; i++ {
//		fmt.Printf("curRow: %v\n", curRow)
//		temp := make([]int64, cLen)
//		for j := 0; j < cLen; j++ {
//			var maxVal int64 = 0
//			for k := 0; k < cLen; k++ {
//				maxVal = max(maxVal, curRow[k]+int64(points[i][j])-int64(absDiff(k, j)))
//			}
//			temp[j] = maxVal
//		}
//		fmt.Printf("Temp: %v\n", temp)
//		for j := 0; j < cLen; j++ {
//			curRow[j] = temp[j]
//		}
//	}
//	res := curRow[0]
//	for i := 1; i < cLen; i++ {
//		if res < curRow[i] {
//			res = curRow[i]
//		}
//	}
//	return res
//}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func absDiff(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func main() {
	points := [][]int{{1, 5}, {3, 2}, {4, 2}}
	fmt.Printf("Res: %v\n", maxPoints(points))
}

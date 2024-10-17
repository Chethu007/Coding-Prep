package main

import (
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	minTime := math.MaxInt32
	n := len(timePoints)
	sort.Strings(timePoints)
	for i := 0; i < n-1; i++ {
		if timePoints[i] == timePoints[i+1] {
			return 0
		}
		//println("Inside the block", i)
		temp1H, _ := strconv.Atoi(timePoints[i][:2])
		temp2H, _ := strconv.Atoi(timePoints[i+1][:2])
		temp1M, _ := strconv.Atoi(timePoints[i][3:5])
		temp2M, _ := strconv.Atoi(timePoints[i+1][3:5])
		//println(temp1H, temp2H, temp1M, temp2M)
		hourDiff := temp2H - temp1H
		minDiff := temp2M + (hourDiff * 60) - temp1M
		minTime = min4(minTime, minDiff)
	}
	//println(minTime)
	firstH, _ := strconv.Atoi(timePoints[0][:2])
	firstM, _ := strconv.Atoi(timePoints[0][3:5])
	lastH, _ := strconv.Atoi(timePoints[n-1][:2])
	lastM, _ := strconv.Atoi(timePoints[n-1][3:5])
	if firstH < 12 && lastH > 12 {
		lastH = 23 - lastH
		lastM = 60 - lastM
		minDiff := firstH*60 + lastH*60 + firstM + lastM
		minTime = min4(minTime, minDiff)
	}
	return minTime
}

func main() {
	strings := []string{"00:00", "22:59", "01:00"}
	println(findMinDifference(strings))
}

func min4(a, b int) int {
	if a < b {
		return a
	}
	return b
}

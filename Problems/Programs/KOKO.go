package main

import (
	"math"
)

//func minEatingSpeed(piles []int, h int) int {
//	//sort.Ints(piles)
//	fmt.Printf("Array: %v\n", piles)
//	max := piles[0]
//	n := len(piles)
//	for i := 0; i < n; i++ {
//		if max < piles[i] {
//			max = piles[i]
//		}
//	}
//	minvalue := max
//	l, r := 1, max
//	for l <= r {
//		mid := (l + r) / 2
//		//println("l, mid and r", l, " ", mid, " ", r)
//		////println("minval ", minvalue)
//		hours := 0
//		for _, val := range piles {
//			hours += val / mid
//			//println("val / mid = hours", val, " ", mid, " ", hours)
//			if val%mid > 0 {
//				hours += 1
//				//println("val % mid = hours", val, " ", mid, " ", hours)
//			}
//		}
//		//println("hours ", hours)
//		if hours <= h {
//			if minvalue > mid {
//				minvalue = mid
//			}
//			//println("minval ", minvalue)
//			r = mid - 1
//		} else {
//			l = mid + 1
//		}
//	}
//	return minvalue
//}

func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	n := len(piles)
	for i := 0; i < n; i++ {
		if max < piles[i] {
			max = piles[i]
		}
	}
	minvalue := max
	l, r := 1, max
	for l <= r {
		mid := (l + r) / 2
		hours := 0
		for _, val := range piles {
			// hours += val / mid
			// if val%mid > 0 {
			// 	hours += 1
			// }
			hours += int(math.Ceil(float64(val) / float64(mid)))
		}
		if hours <= h {
			if minvalue > mid {
				minvalue = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return minvalue
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	res := minEatingSpeed(piles, 6)
	println("Result: ", res)
	println(int(math.Ceil(1.000000001)))
}

package main

func minOperations(nums []int) int {
	count := 0
	for _, val := range nums {
		if count%2 != 0 {
			if val == 0 {
				val = 1
			} else {
				val = 0
			}
		}
		if val%2 != 1 {
			count++
		}
	}
	return count
}

func main() {
	//[1,0,1,1,1,0,1,0,0,0]
	println(minOperations([]int{1, 0, 1, 1, 1, 0, 1, 0, 0, 0}))
}

//func minOperations(nums []int) int {
//	count, z, o := 0, 0, 0
//	for _, val := range nums {
//		if val == 0 {
//			z++
//		} else {
//			o++
//		}
//	}
//	if z > o {
//		for _, val := range nums {
//			if count%2 != 0 {
//				if val == 0 {
//					val = 1
//				} else {
//					val = 0
//				}
//			}
//			if val%2 != 1 {
//				count++
//			}
//		}
//	} else {
//		for _, val := range nums {
//			if count%2 != 0 {
//				if val == 0 {
//					val = 1
//				} else {
//					val = 0
//				}
//			}
//			if val%2 != 0 {
//				count++
//			}
//		}
//		count++
//	}
//	return count
//}

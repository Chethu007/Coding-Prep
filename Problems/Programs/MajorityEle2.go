package main

import "fmt"

func majorityElement(nums []int) []int {
	// Approach 1 - HashMap
	// var res []int
	// m := make(map[int]int)
	// mini := len(nums)/3 + 1
	// for _, num := range nums {
	//     m[num]++
	// }
	// for key, val := range m {
	//     if val >= mini {
	//         res = append(res, key)
	//     }
	// }
	// return res

	// Optimal Solution
	var res []int
	mini := len(nums)/3 + 1
	var ele1, ele2 int
	cnt1, cnt2 := 0, 0
	for i := range nums {
		if cnt1 == 0 && ele2 != nums[i] {
			ele1 = nums[i]
			cnt1 = 1
		} else if cnt2 == 0 && ele1 != nums[i] {
			ele2 = nums[i]
			cnt2 = 1
		} else if ele1 == nums[i] {
			cnt1++
		} else if ele2 == nums[i] {
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}
	println("Mini value", mini)
	println("Ele1 and Ele2:", ele1, " ", ele2)
	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if ele1 == num {
			cnt1++
		} else if ele2 == num {
			cnt2++
		}
	}
	println("Ele1 and Ele2:", cnt1, " ", cnt2)
	if cnt1 >= mini {
		res = append(res, ele1)
	}
	if cnt2 >= mini {
		res = append(res, ele2)
	}

	return res
}

func main() {
	nums := []int{0, 0, 0}
	fmt.Printf("Res: %v", majorityElement(nums))
}

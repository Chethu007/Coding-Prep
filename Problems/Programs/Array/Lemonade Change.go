package main

import "fmt"

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := range bills {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			ten++
			if five == 0 {
				println("1", i)
				return false
			}
			five--
		} else {
			if five == 0 {
				return false
			} else if ten == 0 {
				if five <= 2 {
					return false
				}
				five -= 3
			} else {
				five--
				ten--
			}
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
}

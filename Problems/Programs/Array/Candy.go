package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	temp := make([]int, n)
	for i := range temp {
		temp[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			temp[i] = temp[i-1] + 1
		}
	}
	fmt.Printf("Temp: %v\n: ", temp)
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if temp[i] <= temp[i+1] {
				temp[i] = temp[i+1] + 1
			}
		}
	}
	fmt.Printf("Temp: %v\n: ", temp)
	count := 0
	for i := range temp {
		count += temp[i]
	}
	return count
}

func main() {
	fmt.Println(candy([]int{1, 3, 4, 5, 2}))
}

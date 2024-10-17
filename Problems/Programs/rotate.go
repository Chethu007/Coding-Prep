//package main
//
//import (
//	"math"
//	"sync"
//	"time"
//)
//
////func maxRotateFunction(nums []int) int {
////	max := math.MinInt32
////	for i, _ := range nums {
////		res := Rotateonce(i, nums)
////		if res > max {
////			max = res
////		}
////	}
////	return max
////}
////
////func Rotateonce(k int, a []int) int {
////	n := len(a)
////	print("Original array: ")
////	for i := 0; i < n; i++ {
////		print(a[i], " ")
////	}
////	println()
////	//for j := 0; j < k; j++ {
////	temp := a[n-1]
////	for i := n - 1; i > 0; i-- {
////		a[i] = a[i-1]
////	}
////	a[0] = temp
////	//}
////
////	res := 0
////	for i := 0; i < n; i++ {
////		res += i * a[i]
////		print(a[i], " ")
////	}
////	print("res:", res)
////	println("")
////	return res
////}
//
//func maxRotateFunction(nums []int) int {
//	max := math.MinInt32
//	res := make(chan int)
//	wg := sync.WaitGroup{}
//	for i := 0; i < len(nums); i++ {
//		wg.Add(1)
//		go Rotateonce(nums, res)
//	}
//	func() {
//		for r := range res {
//			if r > max {
//				max = r
//			}
//			wg.Done()
//		}
//	}()
//	wg.Wait()
//	close(res)
//	time.Sleep(2 * time.Second)
//	return max
//}
//
//func Rotateonce(a []int, r chan int) {
//	n := len(a)
//	temp := a[n-1]
//	for i := n - 1; i > 0; i-- {
//		a[i] = a[i-1]
//	}
//	a[0] = temp
//	res := 0
//	for i := 0; i < n; i++ {
//		res += i * a[i]
//		print(a[i], " ")
//	}
//	print("res:", res)
//	println("")
//	r <- res
//}
//
//func main() {
//	nums := []int{4, 3, 2, 6}
//	res := maxRotateFunction(nums)
//	println(res)
//}

package main

import (
	"fmt"
	"math"
	"sync"
)

func maxRotateFunction(nums []int) int {
	max := math.MinInt32
	res := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(nums)) // Update WaitGroup count

	for i, _ := range nums {
		go Rotateonce(i, nums, res, &wg) // Pass WaitGroup as a pointer
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(res) // Close the channel after all results are received
	}()

	for r := range res {
		if r > max {
			max = r
		}
	}

	return max
}

func Rotateonce(k int, nums []int, r chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	a := make([]int, len(nums))
	copy(a, nums) // Signal completion to WaitGroup
	n := len(a)
	for j := 0; j < k; j++ {
		temp := a[n-1]
		for i := n - 1; i > 0; i-- {
			a[i] = a[i-1]
		}
		a[0] = temp
	}
	res := 0
	for i := 0; i < n; i++ {
		res += i * a[i]
		//print(a[i], " ")
	}
	//print("res:", res)
	//println("")
	r <- res
}

func main() {
	nums := []int{4, 3, 2, 6}
	res := maxRotateFunction(nums)
	fmt.Println(res)
}

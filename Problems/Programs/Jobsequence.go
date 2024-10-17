package main

import (
	"fmt"
	"sort"
)

// Job represents a job
type Job struct {
	id     rune // Job Id
	dead   int  // Deadline of job
	profit int  // Profit if job is over before or on deadline
}

//	func printJobScheduling(arr []Job) {
//		// Sort all jobs according to decreasing order of profit
//		sort.Slice(arr, func(i, j int) bool {
//			return arr[i].profit > arr[j].profit
//		})
//
//		n := len(arr)
//		result := make([]int, n) // To store result (Sequence of jobs)
//		slot := make([]bool, n)  // To keep track of free time slots
//
//		// Initialize all slots to be free
//		for i := 0; i < n; i++ {
//			slot[i] = false
//		}
//
//		// Iterate through all given jobs
//		for i := 0; i < n; i++ {
//			// Find a free slot for this job (Note that we start
//			// from the last possible slot)
//			for j := min(n, arr[i].dead) - 1; j >= 0; j-- {
//				// Free slot found
//				if !slot[j] {
//					result[j] = i  // Add this job to result
//					slot[j] = true // Make this slot occupied
//					break
//				}
//			}
//		}
//
//		// Print the result
//		for i := 0; i < n; i++ {
//			if slot[i] {
//				fmt.Printf("%c ", arr[result[i]].id)
//			}
//		}
//	}
func printJobScheduling(arr []Job) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].profit > arr[j].profit
	})

	n := len(arr)
	res := make([]rune, n)
	slot := make([]bool, n)

	for k := 0; k < n; k++ {
		slot[k] = false
	}

	for i := 0; i < n; i++ {
		for j := min(n, arr[i].dead) - 1; j >= 0; j-- {
			if !slot[j] {
				slot[j] = true
				res[j] = arr[i].id
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		if slot[i] {
			fmt.Printf("%c ", res[i])
		}
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	arr := []Job{
		{'a', 1, 100},
		{'b', 4, 19},
		{'c', 2, 47},
		{'d', 3, 22},
		{'e', 4, 25},
	}

	fmt.Println("Following is maximum profit sequence of jobs:")
	printJobScheduling(arr)
}

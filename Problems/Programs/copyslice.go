package main

import "fmt"

func main() {
	// Define the original slices
	n1 := []int{1, 2, 3, 4, 5, 6}
	n2 := []int{2, 3, 4}

	// Copy elements from n2 to n1
	copy(n1, n2)

	// Print n1 after copying
	fmt.Println("n1 after copying n2:", n1)

	n1 = []int{1, 2, 3, 4, 5, 6}
	n2 = []int{2, 3, 4}

	// Create a new slice by slicing n2
	n1 = append([]int(nil), n2...)

	// Print n1 after creating a new slice
	fmt.Println("n1 after creating a new slice from n2:", n1)
}

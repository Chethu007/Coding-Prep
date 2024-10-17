package main

import "fmt"

func main() {
	nums := []int{11, 21, 21, 31, 41, 41}

	// Create a map m where keys may have multiple values
	m := make(map[int][]int)

	// Iterate over nums and populate the map with key-value pairs
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		value := i

		// Check if the key already exists in the map
		if _, ok := m[key]; ok {
			// Key exists, append the value to the existing slice
			m[key] = append(m[key], value)
		} else {
			// Key does not exist, create a new slice with the value
			m[key] = []int{value}
		}
	}

	// Print the map keys and values
	for key, values := range m {
		fmt.Printf("Key: %d, Values: %v\n", key, values)
	}
	for key, values := range m {
		if len(values) == 1 {
			fmt.Printf("Key: %d, Value: %d\n", key, values[0])
		}
	}
}

func singleNumber(nums []int) int {
	var m map[int][]int
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		value := i
		if _, ok := m[key]; ok {
			m[key] = append(m[key], value)
		} else {
			m[key] = []int{value}
		}
	}
	for key, _ := range m {
		if len(m[key]) == 1 {
			return key
		}
	}
	return nums[0]
}

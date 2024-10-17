package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func isUnique(listOfLists [][]int) [][]int {
	// Create a map to store unique lists
	uniqueLists := make(map[string]bool)
	var res [][]int

	// Iterate through each list
	for _, list := range listOfLists {
		// Sort the list to ensure consistent comparison
		sortedList := make([]int, len(list))
		copy(sortedList, list)
		//sortIntSlice(sortedList)
		sort.Ints(sortedList)
		// Convert the sorted list to a string
		listStr := fmt.Sprintf("%v", sortedList)

		// Check if the list has been encountered before
		if uniqueLists[listStr] {
			continue // If encountered before, not unique
		}

		// If not encountered before, mark it as encountered
		uniqueLists[listStr] = true
		//res = append(res, strToInt(listStr))
		res = append(res, sortedList)
	}

	return res
}

func strToInt(str string) []int {
	str = strings.Trim(str, "[]")

	// Split the string by commas
	tokens := strings.Fields(str)

	// Create a slice to store the integers
	intList := make([]int, len(tokens))

	// Convert each token into an integer
	for i, token := range tokens {
		num, err := strconv.Atoi(strings.TrimSpace(token))
		if err != nil {
			fmt.Println("Error:", err)
			return []int{}
		}
		intList[i] = num
	}
	return intList
}

func sortIntSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func main() {
	listOfLists := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{4, 5, 6},
		{1, 2, 3, 4},
		{1, 2, 3},
	}

	fmt.Println("List of lists:", listOfLists)
	fmt.Printf("Unique list: %v", isUnique(listOfLists))
}

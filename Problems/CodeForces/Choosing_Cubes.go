package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func choosingCubes(arr []int, f, k, n int) string {
	//fmt.Printf("Arr: %v", arr)
	k = n - k
	//println("k: ", k, " and n: ", n)
	fav := arr[f-1]
	sort.Ints(arr)
	count := 0
	for i := n - 1; i >= k; i-- {
		if arr[i] == fav {
			count++
		}
	}
	count1 := count
	if count == 0 {
		return "NO"
	} else {
		for i := 0; i < k; i++ {
			if arr[i] == fav {
				count++
			}
		}
	}
	if count != count1 {
		return "MAYBE"
	}
	return "YES"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read the number of test cases
	input, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(input))

	// Loop through each test case
	for i := 0; i < t; i++ {
		// Read n, f, and k
		input, _ := reader.ReadString('\n')
		params := strings.Fields(input)

		n, _ := strconv.Atoi(params[0])
		f, _ := strconv.Atoi(params[1])
		k, _ := strconv.Atoi(params[2])

		// Read the n integers
		input, _ = reader.ReadString('\n')
		strNumbers := strings.Fields(input)

		var numbers []int
		for _, strNum := range strNumbers {
			num, _ := strconv.Atoi(strNum)
			numbers = append(numbers, num)
		}

		// Process the test case
		result := choosingCubes(numbers, f, k, n)
		fmt.Println(result)
	}
}

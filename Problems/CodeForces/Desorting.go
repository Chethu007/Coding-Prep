package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Desorting(arr []int, n int) int {
	minDiff := math.MaxInt32
	for i := 1; i < n; i++ {
		if minDiff > arr[i]-arr[i-1] {
			minDiff = arr[i] - arr[i-1]
		}
	}
	if minDiff < 0 {
		return 0
	}
	return minDiff/2 + 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read the number of test cases
	input, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(input))

	// Loop through each test case
	for i := 0; i < t; i++ {
		// Read n
		input, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(input))

		// Read the n integers
		input, _ = reader.ReadString('\n')
		strNumbers := strings.Fields(input)

		var numbers []int
		for _, strNum := range strNumbers {
			num, _ := strconv.Atoi(strNum)
			numbers = append(numbers, num)
		}

		// Process the test case
		result := Desorting(numbers, n)
		fmt.Println(result)
	}
}

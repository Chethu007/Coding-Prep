package main

import (
	"fmt"
	"math"
)

func main() {
	num := -3
	if num < 0 {
		num = -num // Change the sign by multiplying by -1
	}
	fmt.Println("Changed number:", num)
	num = 0
	fmt.Println("Changed number:", -num)

	fmt.Println("Result", divide(1, -1))
	fmt.Println("Result", divide1(2147483647, -1))
}

func divide(dividend int, divisor int) int {
	flag := 0
	if divisor < 0 {
		flag = 1
		divisor = -divisor
	}
	count := 0
	for dividend >= divisor {
		println("Value before of dividend is", dividend, " Count is,", count)
		dividend -= divisor
		count++
		println("Value after of dividend is", dividend, " Count is,", count)
	}
	if flag == 1 {
		println("Value of count", -count)
		return -count
	}
	println("Value of count", count)
	return count
}

func divide1(dividend int, divisor int) int {
	// Handle overflow cases
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// Determine the sign of the quotient
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	// Convert both dividend and divisor to positive numbers
	dividend = abs(dividend)
	divisor = abs(divisor)

	// Initialize the quotient
	quotient := 0

	// Perform division using repeated subtraction
	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}

	// Apply the sign to the quotient
	quotient *= sign

	return quotient
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

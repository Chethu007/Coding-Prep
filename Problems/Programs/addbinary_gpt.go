package main

import "fmt"

// addBinary adds two binary numbers represented as strings and returns the sum as a string.
func addBinary1(a, b string) string {
	result := ""
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = fmt.Sprintf("%d%s", sum%2, result) // Append the sum modulo 2 to the result string
		carry = sum / 2                             // Calculate the carry for the next iteration
	}

	return result
}

func main() {
	a := "101"
	b := "110"
	sum := addBinary1(a, b)
	fmt.Printf("Sum of %s and %s is %s\n", a, b, sum)
}

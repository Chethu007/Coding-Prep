package main

import (
	"fmt"
	"strings"
)

func main() {
	Num := "12345abc"
	trimLeadingChars := "0"

	result := strings.TrimLeft(Num, trimLeadingChars)
	fmt.Println(result) // Output: "45abc"
}

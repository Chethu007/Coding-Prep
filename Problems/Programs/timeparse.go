package main

import (
	"fmt"
	"time"
)

func main() {
	// Example itemDateStr values
	itemDateStr1 := "15022022"
	itemDateStr2 := "01012023"
	itemDateStr3 := "31072020"
	itemDateStr4 := "25052021"
	itemDateStr5 := "12032025"

	// Parse itemDateStr1 using layout "01022006"
	itemDate1, err := time.Parse("02012006", itemDateStr1)
	if err != nil {
		fmt.Println("Error parsing itemDateStr1:", err)
	} else {
		fmt.Println("Parsed date 1:", itemDate1)
	}

	// Parse itemDateStr2 using layout "01022006"
	itemDate2, err := time.Parse("02012006", itemDateStr2)
	if err != nil {
		fmt.Println("Error parsing itemDateStr2:", err)
	} else {
		fmt.Println("Parsed date 2:", itemDate2)
	}

	// Parse itemDateStr3 using layout "01022006"
	itemDate3, err := time.Parse("02012006", itemDateStr3)
	if err != nil {
		fmt.Println("Error parsing itemDateStr3:", err)
	} else {
		fmt.Println("Parsed date 3:", itemDate3)
	}

	// Parse itemDateStr4 using layout "01022006"
	itemDate4, err := time.Parse("02012006", itemDateStr4)
	if err != nil {
		fmt.Println("Error parsing itemDateStr4:", err)
	} else {
		fmt.Println("Parsed date 4:", itemDate4)
	}

	// Parse itemDateStr5 using layout "01022006"
	itemDate5, err := time.Parse("02012006", itemDateStr5)
	if err != nil {
		fmt.Println("Error parsing itemDateStr5:", err)
	} else {
		fmt.Println("Parsed date 5:", itemDate5)
	}
}

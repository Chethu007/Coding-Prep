package main

import "strings"

func main() {

	sTypeMappingValue := ""
	var sCodes []string

	sCodes = strings.Split(sTypeMappingValue, ",")

	aCodes := []string{"A", "B", "C", "D", "E", "F"}

	accountNumbers := []string{}

	for _, aCode := range aCodes {
		for _, sCode := range sCodes {
			println("sCode: ", sCode)
			for _, accountNumber := range accountNumbers {
				println(aCode, sCode, accountNumber)
			}
		}
	}
}

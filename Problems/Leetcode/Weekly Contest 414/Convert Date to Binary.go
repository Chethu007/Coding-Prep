package main

import (
	"fmt"
	"strconv"
)

func convertDateToBinary(date string) string {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	res, count := "", 0
	temp := []rune{}
	for year != 0 {
		if year%2 == 1 {
			temp = append(temp, '1')
		} else {
			temp = append(temp, '0')
		}
		count++
		year /= 2
	}
	reverse(temp, count)
	res += string(temp) + "-"
	temp = []rune{}
	count = 0
	for month != 0 {
		if month%2 == 1 {
			temp = append(temp, '1')
		} else {
			temp = append(temp, '0')
		}
		count++
		month /= 2
	}
	reverse(temp, count)
	res += string(temp) + "-"
	temp = []rune{}
	count = 0
	for day != 0 {
		if day%2 == 1 {
			temp = append(temp, '1')
		} else {
			temp = append(temp, '0')
		}
		count++
		day /= 2
	}
	reverse(temp, count)
	res += string(temp)
	return res
}

func reverse(temp []rune, count int) {
	for i := 0; i < count/2; i++ {
		temp[i], temp[count-i-1] = temp[count-i-1], temp[i]
	}
}

func main() {
	fmt.Println(convertDateToBinary("1900-01-01"))
}

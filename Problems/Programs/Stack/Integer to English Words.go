package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solve(str string) string {
	m := map[byte]string{
		'1': "One ",
		'2': "Two ",
		'3': "Three ",
		'4': "Four ",
		'5': "Five ",
		'6': "Six ",
		'7': "Seven ",
		'8': "Eight ",
		'9': "Nine ",
		'0': "",
	}
	m1 := map[string]string{
		"2":  "Twenty ",
		"3":  "Thirty ",
		"4":  "Forty ",
		"5":  "Fifty ",
		"6":  "Sixty ",
		"7":  "Seventy ",
		"8":  "Eighty ",
		"9":  "Ninety ",
		"0":  "",
		"10": "Ten ",
		"11": "Eleven ",
		"12": "Twelve ",
		"13": "Thirteen ",
		"14": "Fourteen ",
		"15": "Fifteen ",
		"16": "Sixteen ",
		"17": "Seventeen ",
		"18": "Eighteen ",
		"19": "Nineteen ",
	}
	res := ""
	str = strings.TrimLeft(str, "0")
	n := len(str)
	if n == 3 {
		res += m[str[0]] + "Hundred "
		if str[1] == '1' {
			res += m1[string(str[1])+string(str[2])]
		} else {
			res += m1[string(str[1])] + m[str[2]]
		}

	} else if n == 2 {
		if str[0] == '1' {
			res += m1[string(str[0])+string(str[1])]
		} else {
			res += m1[string(str[0])] + m[str[1]]
		}
	} else if n == 1 {
		res += m[str[0]]
	}
	return res
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	res := ""
	str := strconv.Itoa(num)
	var st []string
	n := len(str)
	for i := n - 1; i >= 0; i -= 3 {
		end := i - 2
		if end < 0 {
			end = 0
		}
		st = append(st, str[end:i+1])
	}
	//fmt.Printf("Stack: %v\n", st)
	stLen := len(st)
	for i := stLen - 1; i >= 0; i-- {
		temp := solve(st[i])
		res += temp
		if temp != "" {
			if i == 5 {
				res += "Quadrillion "
			} else if i == 4 {
				res += "Trillion "
			} else if i == 3 {
				res += "Billion "
			} else if i == 2 {
				res += "Million "
			} else if i == 1 {
				res += "Thousand "
			}
		}
	}
	return res[:len(res)-1]
}

func main() {
	fmt.Println(numberToWords(1000000001))
}

//m := map[string]string{
//	"0":                "Zero",
//	"10":               "Ten",
//	"100":              "One Hundred",
//	"1000":             "One Thousand",
//	"10000":            "Ten Thousand",
//	"100000":           "One Hundred Thousand",
//	"1000000":          "One Million",
//	"10000000":         "Ten Million",
//	"100000000":        "One Hundred Million",
//	"1000000000":       "One Billion",
//	"10000000000":      "Ten Billion",
//	"100000000000":     "One Hundred Billion",
//	"1000000000000":    "One Trillion",
//	"10000000000000":   "Ten Trillion",
//	"100000000000000":  "One Hundred Trillion",
//	"1000000000000000": "One Quadrillion",
//}

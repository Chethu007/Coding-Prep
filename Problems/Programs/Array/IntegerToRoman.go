package main

import "fmt"

//func intToRoman(num int) string {
//	m := map[int]string{
//		1:    "I",
//		5:    "V",
//		10:   "X",
//		50:   "L",
//		100:  "C",
//		500:  "D",
//		1000: "M",
//	}
//	if _, ok := m[num]; ok {
//		return m[num]
//	}
//	rev, count := 0, 0
//	for num != 0 {
//		rem := num % 10
//		rev = rev*10 + rem
//		num = num / 10
//		count++
//	}
//	println(rev)
//	m1 := map[int]string{
//		34: "CD",
//		39: "CM",
//		24: "XL",
//		29: "XC",
//		14: "IV",
//		19: "IX",
//		31: "C",
//		32: "CC",
//		33: "CCC",
//		35: "D",
//		36: "DC",
//		37: "DCC",
//		38: "DCCC",
//		21: "X",
//		22: "XX",
//		23: "XXX",
//		25: "L",
//		26: "LX",
//		27: "LXX",
//		28: "LXXX",
//		11: "I",
//		12: "II",
//		13: "III",
//		15: "V",
//		16: "VI",
//		17: "VII",
//		18: "VIII",
//		41: "M",
//		42: "MM",
//		43: "MMM",
//	}
//	res := ""
//	for i := count; i > 0; i-- {
//		x := i*10 + rev%10
//		println(i, " ", x)
//		res += m1[x]
//		rev /= 10
//	}
//	return res
//}

func intToRoman(num int) string {
	var m = []struct {
		val    int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := ""
	for _, entry := range m {
		//println(entry.val, num)
		for num >= entry.val {
			res += entry.symbol
			num -= entry.val
		}
	}
	return res
}

func main() {
	fmt.Println(intToRoman(3749))
}

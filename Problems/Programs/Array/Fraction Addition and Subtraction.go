package main

import (
	"strconv"
)

func fractionAddition(expression string) string {
	var num, den []int
	var temp, start int
	n := len(expression)

	isDenominator := false
	for start < n {
		if isDenominator {
			i := start
			for i < n && (expression[i] != '+' && expression[i] != '-') {
				i++
			}
			temp, _ = strconv.Atoi(expression[start:i])
			den = append(den, temp)
			start = i
			isDenominator = false
		} else {
			i := start
			for expression[i] != '/' {
				i++
			}
			temp, _ = strconv.Atoi(expression[start:i])
			num = append(num, temp)
			start = i + 1
			isDenominator = true
		}
	}
	var res string
	if len(num) == 1 {
		res = compute(num[0], den[0])
		return res
	}
	//fmt.Printf("Numerator: %v, Denominator: %v\n", num, den)
	lcmVal := lcmOfSlice(den)
	nVal, dVal := 0, lcmVal
	for i := range num {
		nVal += num[i] * lcmVal / den[i]
	}
	//println("nVal: ", nVal, "dVal: ", dVal)
	res = compute(nVal, dVal)
	return res
}

func compute(a, b int) string {
	gcdVal := gcd(a, b)
	//println("gcdVal: ", gcdVal)
	var res string
	if gcdVal == 1 {
		res = strconv.Itoa(a) + "/" + strconv.Itoa(b)
	} else {
		a = a / gcdVal
		b = b / gcdVal
		if a < 0 || b < 0 {
			res = "-" + strconv.Itoa(absVal(a)) + "/" + strconv.Itoa(absVal(b))
		} else {
			res = strconv.Itoa(a) + "/" + strconv.Itoa(b)
		}
	}
	return res
}

func absVal(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func lcmOfSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}

func main() {
	println(fractionAddition("-10/3-5/3-1/4-5/1+7/10"))
}

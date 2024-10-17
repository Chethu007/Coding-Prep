package main

func reverseString(s string) string {
	// Convert the string to a rune slice
	runes := []rune(s)

	// Get the length of the rune slice
	length := len(runes)

	// Iterate over half of the slice and swap characters
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	// Convert the reversed rune slice back to a string
	reversed := string(runes)
	return reversed
}

//func addBinary(a string, b string) string {
//	println("addBinary", a, " ", b)
//	aa := bin(reverseString(a))
//	bb := bin(reverseString(b))
//	println("aa and bb", aa, " ", bb)
//	x := aa + bb
//	println("X value", x)
//	if x == 0 {
//		return "0"
//	}
//	return tobin(x)
//}
//
//func bin(x string) int {
//	println("Bin", x)
//	exponent, res := 0.0, 0.0
//	//for i := 0; i < len(x); i++ {
//	//	println("Value of i:", i, " and value of x[i]:", x[i], " x value is", x)
//	//}
//	for i := 0; i < len(x); i++ {
//		if x[i] == '1' {
//			res += math.Pow(2, exponent)
//		}
//		exponent += 1.0
//	}
//	return int(res)
//}
//
//func tobin(x int) string {
//	println("Tobin", x)
//	var str string
//	for x != 0 {
//		println("Tobin, Value of x", x)
//		rem := x % 2
//		str += strconv.Itoa(rem)
//		x = x / 2
//	}
//	return reverseString(str)
//}

func addBinary(a string, b string) string {
	res, carry := "", 0
	i := len(a)
	j := len(b)
	for i > 0 || j > 0 {
		sum := carry
		if i > 0 {
			if a[i] == '1' {
				sum += 1
			}
			i--
		}
		if j > 0 {
			if a[i] == '1' {
				sum += 1
			}
			j--
		}
		if sum > 1 {
			carry = 1
		}
		res += string(sum % 2)
	}
	if carry == 1 {
		res += string(carry)
	}
	return reverseString(res)
}

func main() {
	a := "11"
	b := "1"
	println(addBinary(a, b))
}

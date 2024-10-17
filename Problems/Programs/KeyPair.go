package main

import (
	"fmt"
)

func main() {
	hash := make(map[int]int)
	var Y, val1, numberToFind int
	var exist bool
	flag := false
	arr := []int{12, 42, 93, 41, 35, 26, 17, 28}
	fmt.Print("Enter the number: ")
	fmt.Scanln(&Y)
	for i, num := range arr {
		numberToFind = Y - num
		val1, exist = hash[numberToFind]
		println(num, val1, i)
		if exist {
			flag = true
			println("Found:", val1, "-->", hash[numberToFind], "numberToFind", numberToFind)
			break
		} else {
			hash[num] = i
		}
	}
	for key, val := range hash {
		println(key, "-->", val)
	}
	for _, val := range hash {
		if val1 == val {
			println("Index:", hash[Y-val], hash[numberToFind])
			println("Value:", (Y - numberToFind), "+", numberToFind)
		}
	}
	if !flag {
		println("Not found")
	}

}

//if _, found := hash[numberToFind]; found {
//return true
//}
//hash[num] = i
//}

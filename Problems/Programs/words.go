package main

import "strings"

func main() {
	text := "I am Chethan . My name is Chethan . Chethan is my name . Name is Chethan"
	Slice := strings.Fields(text)
	res := map[string]int{}
	for _, word := range Slice {
		res[strings.ToLower(word)]++
	}
	for key, value := range res {
		println(key, "-->", value)
	}
}

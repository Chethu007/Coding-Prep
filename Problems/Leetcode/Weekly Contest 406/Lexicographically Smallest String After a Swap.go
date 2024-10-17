package main

import "fmt"

func getSmallestString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		first, second := int(r[i]-'0'), int(r[i+1]-'0')
		if first%2 == second%2 {
			if first > second {
				r[i], r[i+1] = r[i+1], r[i]
				break
			}
		}
	}
	return string(r)
}

func main() {
	fmt.Println(getSmallestString("absds45320"))
}

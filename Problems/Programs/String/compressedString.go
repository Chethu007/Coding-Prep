package main

import "strconv"

//func compressedString(word string) string {
//	n := len(word)
//	comp := ""
//	i := 0
//	for i < n {
//		maxLen := 1
//		for i < n-1 && maxLen < 9 && word[i] == word[i+1] {
//			maxLen++
//			i++
//		}
//		comp += strconv.Itoa(maxLen) + string(word[i])
//		i++
//	}
//	return comp
//}

func compressedString(word string) string {
	n := len(word)
	comp := ""
	maxLen, count := 1, 0
	for i := 0; i < n-1; i++ {
		if word[i] == word[i+1] && maxLen < 9 {
			maxLen++
			continue
		} else {
			comp += strconv.Itoa(maxLen) + string(word[i])
			count += maxLen
			maxLen = 1
		}
	}
	if count != n {
		comp += strconv.Itoa(maxLen) + string(word[n-1])
	}
	return comp
}

func main() {
	word := "abcde"
	println("Res: ", compressedString(word), " ", len(compressedString(word)))
}

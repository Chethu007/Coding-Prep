package main

//func numberOfSpecialChars(word string) int {
//	m := make(map[rune]bool)
//	r := []rune(word)
//	n := len(r)
//	for i := 0; i < n-1; i++ {
//		if isUpper(r[i]) {
//			continue
//		}
//		iC, iS := math.MaxInt32, 0
//		for j := i + 1; j < n; j++ {
//			if r[j] == r[i] {
//				iS = j
//				continue
//			}
//			if r[j] == r[i]-32 {
//				iC = min1(iC, j)
//			}
//		}
//		if iC > iS && iC != math.MaxInt32 {
//			m[r[i]] = true
//		}
//	}
//	for key, val := range m {
//		fmt.Printf("key: %c and val: %v\n", key, val)
//	}
//	return len(m)
//}

func numberOfSpecialChars(word string) int {
	mS := make(map[rune]int)
	mC := make(map[rune]int)
	r := []rune(word)
	n := len(r)
	count := 0
	for i := 0; i < n; i++ {
		if isUpper(r[i]) {
			if _, ok := mC[r[i]]; !ok {
				mC[r[i]] = i
			}
		} else {
			mS[r[i]] = i
		}
	}
	//for key, val := range mS {
	//	fmt.Printf("key1: %c and val1: %v\n", key, val)
	//}
	//for key, val := range mC {
	//	fmt.Printf("key2: %c and val2: %v\n", key, val)
	//}
	for key1, _ := range mC {
		//fmt.Printf("mS[%c]: %c and mC[%c]: %c\n", mS[key1], key1, mC[key1], key1+32)
		if _, ok := mS[key1+32]; ok {
			//fmt.Printf("mS[%c]: %d and mC[%c]: %d\n", key1+32, mS[key1+32], key1, mC[key1])
			if mS[key1+32] < mC[key1] {
				count++
				println("Count", count)
			}
		}
	}
	return count
}

func min1(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func isUpper(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func main() {
	word := "cCcdDC"
	println("Res: ", numberOfSpecialChars(word))
}

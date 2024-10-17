package main

import "math"

func minWindow(s string, t string) string {
	window := make(map[rune]int)
	need := make(map[rune]int)
	l, r, start, match := 0, 0, 0, 0
	minLen := math.MaxInt32
	for _, val := range t {
		need[val]++
	}
	for r < len(s) {
		c1 := rune(s[r])
		_, ok := need[c1]
		if ok {
			window[c1]++
			println("window[c1]:", window[c1], " need[c1]:", need[c1])
			if window[c1] == need[c1] {
				match++
			}
		}
		println("Match:", match, " minlen:", minLen, " lenofT:", len(t))

		r++
		for match == len(t) {
			println("value of r", r)
			if r-l < minLen {
				// Updates the position and length of the smallest string
				start = l
				minLen = r - l
			}
			c2 := rune(s[l])
			_, ok := need[c2]
			if ok {
				window[c2]--
				if window[c2] < need[c2] {
					match--
				}
			}
			l++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+minLen]
	}
}

func main() {
	println(minWindow("aa", "aa"))
}

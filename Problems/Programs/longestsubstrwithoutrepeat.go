package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l, r, maxlen := 0, 0, 0
	sr := []rune(s)
	//maxstr := ""
	m := map[rune]int{}
	for l < len(sr) && r < len(sr) {
		ch := sr[r]
		fmt.Printf("char at r is %s-%v and char at l %s-%v, length of map %d\n", string(ch), r, string(sr[l]), l, len(m))
		_, ok := m[ch]
		if !ok {
			m[ch]++
			r++
			//println("Length of r", r)
			if maxlen < r-l {
				maxlen = r - l
				//maxstr = string(sr[l:r-1])
			}
		} else {
			m[sr[l]]--
			l++
			//println("Length of l", l)
		}
	}
	return maxlen
}

func main() {
	s := "pwwkew"
	println(lengthOfLongestSubstring(s))
}

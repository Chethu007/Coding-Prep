package main

import "sort"

type kv struct {
	Key byte
	Val int
}

func minimumPushes(word string) int {
	count := 0
	m := make(map[byte]int)
	for i := range word {
		m[word[i]]++
	}
	var kvslice []kv
	for key, val := range m {
		kvslice = append(kvslice, kv{key, val})
	}
	sort.Slice(kvslice, func(i, j int) bool {
		return kvslice[i].Val > kvslice[j].Val
	})
	for i, val := range kvslice {
		if i >= 0 && i < 8 {
			count += val.Val
		} else if i >= 8 && i < 16 {
			count += 2 * val.Val
		} else if i >= 16 && i < 24 {
			count += 3 * val.Val
		} else {
			count += 4 * val.Val
		}
	}
	return count
}

func main() {
	println(minimumPushes("aabbccddeeffgghhiiiiii"))
}

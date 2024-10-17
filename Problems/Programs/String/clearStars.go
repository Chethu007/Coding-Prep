package main

func clearStars(s string) string {
	r := []byte{}
	res := ""
	for i := range s {
		if s[i] == '*' {
			r = []byte(smallestString(string(r)))
		} else {
			r = append(r, s[i])
		}
		res = string(r)
	}
	return res
}

func smallestString(s string) string {
	minValue := 100
	var minChar byte
	for i := range s {
		temp := int(s[i] - 'a')
		if minValue > temp {
			minValue = temp
			minChar = s[i]
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == minChar {
			return s[:i] + s[i+1:]
		}
	}
	return s
}

func main() {
	word := "aaba"
	println("Res: ", clearStars(word))
}

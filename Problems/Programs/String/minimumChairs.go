package main

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minimumChairs(s string) int {
	count, minChair := 0, -1
	for i := range s {
		if s[i] == 'E' {
			count++
			minChair = Max(minChair, count)
		} else {
			count--
		}
	}
	return minChair
}

func main() {
	word := "ELEELEELLL"
	println("Res: ", minimumChairs(word))
}

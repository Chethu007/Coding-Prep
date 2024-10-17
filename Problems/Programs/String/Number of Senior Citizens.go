package main

func countSeniors(details []string) int {
	count := 0
	for _, s := range details {
		//println(int('6'), " --> ", int(s[10]), " --> ", s[11], " --> ", int(s[11]))
		if int(s[11]) < int('6') {
			continue
		} else if int(s[12]) == int('0') {
			continue
		}
		count += 1
	}
	return count
}

func main() {
	println(countSeniors([]string{"1313579440F2036", "2921522980M5644"}))
}

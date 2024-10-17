package main

func losingPlayer(x int, y int) string {
	if x*4 <= y {
		if x%2 == 0 {
			return "Bob"
		} else {
			return "Alice"
		}
	} else {
		if (y/4)%2 == 0 {
			return "Bob"
		} else {
			return "Alice"
		}
	}
}

func main() {
	println(losingPlayer(4, 11))
}

package main

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	x1, y1, x2, y2 := int(coordinate1[0]), int(coordinate1[1]), int(coordinate2[0]), int(coordinate2[1])
	if x1%2 == x2%2 {
		if y1%2 == y2%2 {
			return true
		} else {
			return false
		}
	} else {
		if y1%2 == y2%2 {
			return false
		} else {
			return true
		}
	}
}

func main() {
	println("Res", checkTwoChessboards("a1", "h3"))
}

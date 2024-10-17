package main

func main() {

	arr := []int{1, 2, 3}
	arr1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	x := 5
cycleLoop:
	for _, ele := range arr {
		x++
		println("First loop:", ele, "X value", x)
		println("Second loop starts:")
		for _, ele1 := range arr1 {
			if ele1 == x {
				continue
			}

			if ele1 == 4 {
				println()
				continue cycleLoop
			}
			print(ele1, " ")
		}
		println()
	}
}

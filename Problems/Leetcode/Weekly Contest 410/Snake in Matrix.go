package main

import "fmt"

func finalPositionOfSnake(n int, commands []string) int {
	i, j := 0, 0
	for _, val := range commands {
		if val == "UP" {
			i--
		} else if val == "DOWN" {
			i++
		} else if val == "RIGHT" {
			j++
		} else {
			j--
		}
	}
	return (i * n) + j
}

func main() {
	fmt.Printf("Res :%v\n", finalPositionOfSnake(3, []string{"DOWN", "RIGHT", "UP"}))
}

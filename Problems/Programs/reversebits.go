package main

import (
	"fmt"
)

func reverseBits(num uint) uint {
	var result uint
	numBits := bitLength(num)

	for i := 0; i < numBits; i++ {
		result <<= 1      // Left shift the result by 1 position
		result |= num & 1 // Set the least significant bit of result with the current bit of num
		num >>= 1         // Right shift num by 1 position to process the next bit
	}
	return result
}

// Function to calculate the number of bits needed for a given unsigned integer
func bitLength(num uint) int {
	if num == 0 {
		return 1 // 0 needs at least 1 bit to represent
	}
	count := 0
	for num > 0 {
		count++
		num >>= 1
	}
	return count
}

func main() {
	num := uint(0) // Example input
	reversed := reverseBits(num)
	fmt.Printf("Original: %b\nReversed: %b\n", num, reversed)
}

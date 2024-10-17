package main

import "fmt"

func main() {
	// Setting a bit
	var flag uint8
	fmt.Printf("Flag: %b", flag)
	println()
	flag |= (1 << 2)                                // Set the 2nd bit (zero-based index) of flag
	fmt.Printf("After setting bit 2: %08b\n", flag) // Output: 00000100

	fmt.Printf("Flag: %b", flag)
	println()
	// Clearing a bit
	flag &^= (1 << 3)                                // Clear the 3rd bit (zero-based index) of flag
	fmt.Printf("After clearing bit 3: %08b\n", flag) // Output: 00000000

	fmt.Printf("Flag: %b", flag)
	println()
	// Toggling a bit
	flag ^= (1 << 1)                                 // Toggle the 1st bit (zero-based index) of flag
	fmt.Printf("After toggling bit 1: %08b\n", flag) // Output: 00000010

	fmt.Printf("Flag: %b", flag)
	println()
	// Checking if a bit is set
	var num uint8 = 0b1101_0010                // Binary: 11010010
	bitIsSet := num&(1<<3) != 0                // Check if the 4th bit (zero-based index) is set in num
	fmt.Printf("Is bit 4 set? %t\n", bitIsSet) // Output: Is bit 4 set? true

	fmt.Printf("Flag: %b", flag)
	println()
	// Extracting bits
	extractedBits := (num >> 3) & 0b0000_1111               // Extract bits 3 to 6 (zero-based) from num
	fmt.Printf("Extracted bits 3-6: %04b\n", extractedBits) // Output: 0100
}

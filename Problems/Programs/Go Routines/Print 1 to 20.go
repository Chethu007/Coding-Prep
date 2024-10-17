//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	// WaitGroup to ensure main waits for both goroutines to finish
//	var wg sync.WaitGroup
//	wg.Add(2) // Two goroutines
//
//	// Create channels to signal between goroutines
//	oddChan := make(chan bool, 1)  // Buffered channel to avoid deadlock
//	evenChan := make(chan bool, 1) // Buffered channel to avoid deadlock
//
//	// Goroutine to print odd numbers
//	go func() {
//		defer wg.Done() // Notify that this goroutine is done
//		for i := 1; i <= 19; i += 2 {
//			<-oddChan        // Wait for the signal to print
//			fmt.Println(i)   // Print the odd number
//			evenChan <- true // Signal the even goroutine
//		}
//	}()
//
//	// Goroutine to print even numbers
//	go func() {
//		defer wg.Done() // Notify that this goroutine is done
//		for i := 2; i <= 20; i += 2 {
//			<-evenChan     // Wait for the signal to print
//			fmt.Println(i) // Print the even number
//			if i < 20 {    // Only signal back if it's not the last number
//				oddChan <- true // Signal the odd goroutine
//			}
//		}
//	}()
//
//	// Start the chain by signaling the odd goroutine first
//	oddChan <- true
//
//	// Wait for both goroutines to finish
//	wg.Wait()
//}

package main

import (
	"fmt"
)

func main() {
	// Create channels to signal between goroutines
	oddChan := make(chan bool, 1)  // Buffered channel to avoid deadlock
	evenChan := make(chan bool, 1) // Buffered channel to avoid deadlock
	doneChan := make(chan bool)    // Channel to signal the completion of both goroutines

	// Goroutine to print odd numbers
	go func() {
		for i := 1; i <= 19; i += 2 {
			<-oddChan        // Wait for the signal to print
			fmt.Println(i)   // Print the odd number
			evenChan <- true // Signal the even goroutine
		}
		//doneChan <- true // Signal completion of the odd goroutine
	}()

	// Goroutine to print even numbers
	go func() {
		for i := 2; i <= 20; i += 2 {
			<-evenChan     // Wait for the signal to print
			fmt.Println(i) // Print the even number
			if i < 20 {    // Only signal back if it's not the last number
				oddChan <- true // Signal the odd goroutine
			}
		}
		doneChan <- true // Signal completion of the even goroutine
	}()

	// Start the chain by signaling the odd goroutine first
	oddChan <- true

	// Wait for both goroutines to signal completion
	<-doneChan
}

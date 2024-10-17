package main

import (
	"fmt"
	"time"
)

// Makereq simulates making a request
func Makereq(req string, ch chan<- string) {
	// Simulate some processing time
	time.Sleep(2 * time.Second)
	// Send the response on the channel
	ch <- req + " Response"
}

// process handles the request and response
func process(req string, done chan<- bool, requeue chan<- string) {
	fmt.Println("Processing request:", req)
	// Simulate processing time
	time.Sleep(2 * time.Second)
	fmt.Println("Completed processing:", req)
	// Signal completion
	done <- true
	// Simulate requeueing
	requeue <- req
}

// processreq processes requests and responses
func processreq(req string) {
	hmc := make(chan string, 256)
	requeue := make(chan string, 256)
	done := make(chan bool, 256)

	reqSent := len(req)
	go Makereq(req, hmc)
	completed := 0

handleResponse:
	for {
		select {
		case resp := <-hmc:
			go process(resp, done, requeue)
			continue
		case hm := <-requeue:
			reqSent++
			go Makereq(hm, hmc)
		case <-done:
			completed++
			if reqSent == completed {
				close(done)
				close(hmc)
				break handleResponse
			}
		case <-time.After(10000 * time.Millisecond):
			fmt.Println("Timeout occurred")
			continue
		}
	}
}

func main() {
	go processreq("Request1")
	go processreq("Request2")
	time.Sleep(20 * time.Second) // Simulate main function running for 20 seconds
}

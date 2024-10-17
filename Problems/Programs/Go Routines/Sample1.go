package main

//
//import (
//	"fmt"
//	"strings"
//	"sync"
//)
//
//// Example function signature for the Go routine
//func fetchData(str string) ([]string, error) {
//	// Simulating some work and potential error
//	if str == "error" {
//		return nil, fmt.Errorf("an error occurred for: %s", str)
//	}
//	return []string{str, str + "_result"}, nil
//}
//
//func main() {
//	searchKeywords := []string{"keyword1", "keyword2", "error1", "keyword3"}
//	var cycles []string
//
//	// Channels to collect results and errors
//	resultCh := make(chan []string)
//	errCh := make(chan error)
//
//	// Use a WaitGroup to wait for all goroutines to finish
//	var wg sync.WaitGroup
//
//	// Loop through searchKeywords and start a goroutine for each
//	for _, searchKeyword := range searchKeywords {
//		searchKeyword := searchKeyword // Shadow variable to capture loop variable correctly
//		if strings.TrimSpace(searchKeyword) != "" {
//			wg.Add(1)
//			go func(keyword string) {
//				defer wg.Done()
//				cycle, err := fetchData(keyword)
//				if err != nil {
//					errCh <- err
//					return
//				}
//				resultCh <- cycle
//			}(searchKeyword)
//		}
//	}
//
//	// Close the result and error channels after all goroutines are done
//	go func() {
//		wg.Wait()
//		close(resultCh)
//		close(errCh)
//	}()
//	var err error
//
//	for resultCh != nil {
//		select {
//		case result, ok := <-resultCh:
//			if !ok {
//				println("resultCh closed")
//				resultCh = nil
//			} else {
//				println("resultCh still open")
//				cycles = append(cycles, result...)
//			}
//		case err = <-errCh:
//			if err != nil {
//				fmt.Printf("Error: %v\n", err)
//				return
//			}
//		}
//	}
//
//	fmt.Println("Final cycles:", cycles)
//}

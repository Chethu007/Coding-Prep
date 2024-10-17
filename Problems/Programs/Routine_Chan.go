package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func work(t time.Duration, wg *sync.WaitGroup, s string, res chan string) {
	fmt.Printf("Working %s\n", s)
	time.Sleep(t)
	fmt.Printf("Done %s\n", s)
	res <- fmt.Sprintf("Res %s with work: %d", s, rand.Intn(100))
	wg.Done()
}

func work1(t time.Duration, s string, res chan string) {
	fmt.Printf("Working %s\n", s)
	time.Sleep(t)
	fmt.Printf("Done %s\n", s)
	res <- fmt.Sprintf("Res %s with work: %d", s, rand.Intn(100))
}
func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	res := make(chan string)
	wg.Add(3)
	go work(time.Second*2, &wg, "1", res)
	go work(time.Second*4, &wg, "2", res)
	go work(time.Second*6, &wg, "3", res)
	//go work1(time.Second*2, "1", res)
	//go work1(time.Second*4, "2", res)
	//go work1(time.Second*6, "3", res)
	//wg.Wait()
	//fmt.Printf("Res1: %v", <-res)
	//fmt.Printf("Res2: %v", <-res)
	//for r := range res {
	//	println(r)
	//}

	go func() {
		for r := range res {
			println(r)
		}
		fmt.Printf("Completed in: %v", time.Since(start))
	}()
	wg.Wait()
	close(res)
	time.Sleep(time.Second)

	//println(<-res)
	//println(<-res)
	//println(<-res)
}

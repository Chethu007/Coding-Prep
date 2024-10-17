package main

import (
	"fmt"
	"math"
	"sync"
)

// Fibonaccharsis function checks the validity of the sequence
func fibonaccharsis(i, n, k int64) bool {
	var j int64
	valid := true
	first := n
	second := i
	for j = 2; j < k; j++ {
		temp := second
		second = first - second
		first = temp
		if first < second || second < 0 || first < 0 {
			valid = false
			break
		}
	}
	return valid
}

// Worker function to process a range of values
func worker(n, k int64, jobs <-chan int64, results chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		if fibonaccharsis(i, n, k) {
			results <- 1
		} else {
			results <- 0
		}
	}
}

func main() {
	var t, n, k int64
	fmt.Scan(&t)

	for t > 0 {
		fmt.Scan(&n, &k)
		if (k >= 7 && k >= n) || k > 30 {
			fmt.Println(0)
		} else {
			jobs := make(chan int64, n)
			results := make(chan int64, n)
			var wg sync.WaitGroup

			// Start a fixed number of workers
			numWorkers := 8
			for w := 0; w < numWorkers; w++ {
				wg.Add(1)
				go worker(n, k, jobs, results, &wg)
			}

			// Send jobs to workers
			go func() {
				for i := int64(math.Ceil(float64(n) / 2)); i <= n; i++ {
					jobs <- i
				}
				close(jobs)
			}()

			// Wait for all workers to finish
			go func() {
				wg.Wait()
				close(results)
			}()

			// Collect results
			var res int64 = 0
			for result := range results {
				res += result
			}

			fmt.Println(res)
		}
		t--
	}
}

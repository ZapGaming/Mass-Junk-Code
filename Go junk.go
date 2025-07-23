package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Thread-safe memoization cache using sync.Map
var fibCache sync.Map

func simulateWorkGo(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func fibonacciGo(n int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done() // Ensure wg.Done is called when the function returns

	if n < 0 {
		fmt.Println("Go: Input must be a non-negative integer.")
		results <- -1 // Signal error
		return
	}
	if n == 0 {
		results <- 0
		return
	}
	if n == 1 {
		results <- 1
		return
	}

	// Check cache
	if val, ok := fibCache.Load(n); ok {
		results <- val.(int) // Type assertion required for sync.Map values
		return
	}

	// Simulate work
	simulateWorkGo(1)

	// For truly concurrent calculation without blocking this goroutine waiting for child goroutines,
	// we'd ideally use channels or sync primitives. For simplicity and to demonstrate cache interaction,
	// we will calculate these sequentially within this fibonacciGo call. A more advanced Go pattern
	// would involve managing child tasks via their own goroutines and communicating results back.

	// Using recursion here, but for large N, iterative is better to avoid stack overflow.
	// Here, we use helper channels to get results from recursive calls.
	ch1 := make(chan int)
	ch2 := make(chan int)

	var innerWg sync.WaitGroup
	innerWg.Add(2)

	go fibonacciGo(n-1, &innerWg, ch1)
	go fibonacciGo(n-2, &innerWg, ch2)

	innerWg.Wait() // Wait for both sub-computations to finish
	close(ch1)     // Close channels to allow range loops to terminate
	close(ch2)

	res1 := <-ch1
	res2 := <-ch2
    
	// Close the channels passed from above after we are done with them
    // close(ch1) // Removed; channels shouldn't be closed by receiver unless it's the sole intended reader.
    // close(ch2) // If multiple goroutines are reading from a channel, closing it causes panic.

	// If either sub-computation failed, propagate error.
	if res1 == -1 || res2 == -1 {
		results <- -1
		return
	}

	result := res1 + res2
	fibCache.Store(n, result) // Store in cache
	results <- result
}

func calculateConcurrentFibonacciGo(maxN, numThreads int) {
	fmt.Println("\n--- Go Example ---")
	fmt.Printf("Go: Calculating Fibonacci numbers up to %d concurrently using %d goroutines...\n", maxN, numThreads)
	start := time.Now()

	var wg sync.WaitGroup
	resultsChan := make(chan int, maxN+1) // Buffered channel for results

	// Reset cache for this run (or assume fresh start for demonstration)
	fibCache = sync.Map{} 

	wg.Add(maxN + 1)
	for i := 0; i <= maxN; i++ {
		go fibonacciGo(i, &wg, resultsChan)
	}

	// Wait for all goroutines to signal completion
	wg.Wait()
	close(resultsChan) // Close the channel to signal that no more results will be sent

	var collectedResults []int
	for res := range resultsChan {
		collectedResults = append(collectedResults, res)
	}

	// Check if any errors occurred (indicated by -1)
	hasError := false
	for _, r := range collectedResults {
		if r == -1 {
			hasError = true
			break
		}
	}

	elapsed := time.Since(start)
	if !hasError {
		fmt.Printf("Go: Total time taken for Fibonacci up to %d: %v\n", maxN, elapsed)
	} else {
		fmt.Println("Go: Calculation completed with errors.")
	}
	fmt.Println("Go calculation complete.\n")
}

// Example call: calculateConcurrentFibonacciGo(15, 4);

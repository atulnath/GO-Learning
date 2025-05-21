package main

import (
	"fmt"
	"sync"
	"time"
)

func employee(employ int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when the goroutine completes
	fmt.Printf("Employee %d is working\n", employ)
	// Simulate some work with a sleep
	time.Sleep(time.Second * 2)
	fmt.Printf("Employee %d has finished working\n", employ)
	// Simulate some more work with a sleep

}

func main() {
	//lets create a wait group to wait for the goroutine to finish

	var waitGroup sync.WaitGroup
	for employ := 1; employ <= 3; employ++ {
		waitGroup.Add(1)                // Add a count to the wait group
		go employee(employ, &waitGroup) // Start the goroutine
	}
	waitGroup.Wait() // Wait for all goroutines to finish
	fmt.Println("All employees have finished their work")
}

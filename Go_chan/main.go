package main

import "fmt"

// Step 1: Create a channel using the `make` function
// ch := make(chan string)

// send data

// ch <- "Hello, Channel!"

// receive data

// msg := <-ch

func main() {
	ch := make(chan string)

	// Sending a message to the channel
	go func() {
		ch <- "Running a goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)

	bufferedChannel()
	sendingThreeValues()
}

func bufferedChannel() {
	ch := make(chan string, 2) // capacity of 2

	// Sending data to the buffered channel
	ch <- "Message 1"
	ch <- "Message 2"

	// Receiving data from the buffered channel
	msg1 := <-ch
	msg2 := <-ch

	fmt.Println(msg1)
	fmt.Println(msg2)
}

func sendingThreeValues() {
	ch := make(chan int)

	// Sending three values to the channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
	}()

	for nums := range ch {
		fmt.Println("Received:", nums)
	}
}

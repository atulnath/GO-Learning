package main

import (
	"fmt"
	"time"
)

func sendMessage(channel chan string) {
	time.Sleep(2 * time.Second) // Simulate a delay
	channel <- "Hello, World!"  // Send a message to the channel
}

func main() {
	channels := make(chan string) // Create a channel to communicate between goroutines
	go sendMessage(channels)      // Start the sendMessage goroutine
	meg := <-channels             // Receive the message from the channel
	fmt.Println(meg)              // Print the message received from the channel
}

// print the message<- mag <-channel <- (msg)string

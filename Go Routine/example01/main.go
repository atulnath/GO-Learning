package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("This function is running in a goroutine")
}

func main() {
	fmt.Println("Main function runs before go routine is running")
	go sayHello()               // Start a goroutine
	time.Sleep(3 * time.Second) // Wait for the goroutine to finish
	fmt.Println("Main function is running")
}

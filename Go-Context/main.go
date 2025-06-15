package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	// Start a task that respects the context
	go performTask(ctx)

	// Wait for the task to complete or timeout
	select {
	case <-ctx.Done(): // Cancelled or timed out
		// The context has been cancelled or timed out
		fmt.Println("Main function finished:", ctx.Err())
	case <-time.After(2 * time.Second):
		fmt.Println("Main function finished after 2 seconds")
	}
}

// This function demonstrates how to use context
// to manage cancellation of a task.
func performTask(ctx context.Context) {
	// Simulate a task that respects the context
	select {
	// Check if the context has been cancelled
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
		return
	default:
		// Perform the task
		fmt.Println("Performing task...")

		// Simulate a delay
		time.Sleep(500 * time.Second)

	}
}

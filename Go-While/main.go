package main

import "fmt"

// ulike other languages, Go does not have a while loop.
// Instead, you can use a for loop with a condition to achieve similar functionality.
func main() {

	num := 0
	for num < 5 {
		fmt.Print("Number: ", num, "\n")
		num++
	}

	doWhile()
}

func doWhile() {
	// A do-while loop can be simulated using a for loop with a break condition.
	num := 0
	for {
		fmt.Print("Number: ", num, "\n")
		num++
		if num >= 5 {
			break
		}
	}
}

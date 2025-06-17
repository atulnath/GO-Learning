package main

import "fmt"

func main() {
	rangeForSlice()
	rangeForString()
	rangeForMap()
}

func rangeForMap() {
	students := map[string]int{"Alice": 20, "Bob": 22, "Charlie": 23}
	for name := range students {
		fmt.Printf("Name: %s\n", name)
	}
}
func rangeForString() {
	str := "Hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
func rangeForSlice() {
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

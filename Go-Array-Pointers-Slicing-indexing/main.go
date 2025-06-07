package main

import "fmt"

func main() {

	//Array Declaration
	var arr [5]int
	// Expected Output: [0 0 0 0 0]
	fmt.Println(arr)

	var nums [3]int = [3]int{1, 2, 3}
	// Expected Output: [1 2 3]
	fmt.Println(nums)

	//This function demonstrates how to use pointers with arrays in Go
	// We will create an array using a pointer and then modify its elements.
	ArrayUsingPointers()

	// This function demonstrates how to create and use slices in Go.
	SlicingInGo()

	// This function demonstrates how to create a slice from an array in Go.
	SlicingInGoUsingArray()

	// This function demonstrates how to  use indexing in Go.
	indexingInGo()

	// This function demonstrates how to  use slicing of array elements in Go.
	slicingOfElementInArray()

}

func ArrayUsingPointers() {

	xp := new([5]int) //[5]int is just like what we have seen before
	(*xp)[0] = 10
	(*xp)[1] = 20
	(*xp)[2] = 30
	(*xp)[3] = 40
	(*xp)[4] = 50

	fmt.Println(*xp) // Expected Output: [10 20 30 40 50]
}

func SlicingInGo() {
	x := make([]int, 3, 5) // Create a slice of integers with length 3 and capacity 5
	x[0] = 1
	x[1] = 2
	x[2] = 3
	fmt.Println(x)      // Expected Output: [1 2 3]
	fmt.Println(len(x)) // Expected Output: 3

	//There is two things here, one is length of Slice Array and another is capcity
	fmt.Printf("Length of Array is: %d and the capacity is: %d\n", len(x), cap(x))
}

func SlicingInGoUsingArray() {
	xp := new([5]int)[0:3] // Create a slice from an array of integers with length 3 and capacity 5
	xp[0] = 1
	xp[1] = 2
	xp[2] = 3
	fmt.Println(xp) // Expected Output: [1 2 3]
}

func indexingInGo() {

	x := [4]int{10, 20, 30, 40}

	fmt.Println(x[2], x[3])
}

func slicingOfElementInArray() {
	a := [5]int{2, 3, 4, 5, 6}

	slices := a[0:3]

	//Even we can make slice of slices

	slices2 := slices[1:3]
	fmt.Println(slices)
	fmt.Println(slices2)

	//After slice we can change the value of sliceing array

	slices[1] = 100
	fmt.Println(slices)

}

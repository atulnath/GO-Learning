package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5} // This is an array

	fmt.Println(nums)

	nums2 := []int{10, 20, 30, 40} // This is a slice
	fmt.Println(nums2)

	// Creating a slice from an array
	sliceFromArray := nums[1:4] // This will take elements from index

	fmt.Println(sliceFromArray)

	slices := nums[1:3]
	fmt.Println(slices)

	slice2 := nums2[1:2]

	fmt.Println(slice2)

	// Slice Functions (1) - Len
	fmt.Println("Length of nums2 slice:", len(nums2))
	fmt.Println("Length of sliceFromArray slice:", len(sliceFromArray))

	// Slice Functions (2) - Cap
	fmt.Println("Capacity of nums2 slice:", cap(nums2))
}

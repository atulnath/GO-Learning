package main

import (
	"fmt"
	"net"
)

func main() {
	mask := net.IPv4Mask(255, 255, 255, 0) // Create a subnet mask for IPv4
	fmt.Println("Subnet Mask:", mask)

	num_of_ones, num_of_bits := mask.Size() // Get the number of ones and bits in the mask
	fmt.Println("Number of ones:", num_of_ones)
	fmt.Println("Number of bits:", num_of_bits)
}

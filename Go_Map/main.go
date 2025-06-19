package main

import "fmt"

func main() {
	itemsPrices := map[string]float64{
		"apple":  0.99,
		"banana": 0.59,
		"orange": 0.79,
		"grape":  2.49,
	}

	fmt.Println("Items and their prices \n", itemsPrices)

	for item, price := range itemsPrices {
		fmt.Printf("The price of %s is $%.2f\n", item, price)
	}

	// acessing a specific item
	// If the key does not exist, it will return the zero value for the value type
	fmt.Println("Apple price:", itemsPrices["apple"])
	fmt.Println("Mango price:", itemsPrices["mango"]) // This will return 0.0 since "mango" does not exist

	// lets change the price of an item
	itemsPrices["banana"] = 0.65
	fmt.Println("Updated banana price:", itemsPrices["banana"])

	itemsPrices["kiwi"] = 1.29 // Adding a new item
	fmt.Println("Added kiwi:", itemsPrices["kiwi"])

	delete(itemsPrices, "orange") // Removing an item
	fmt.Println("After deleting orange:", itemsPrices)
}

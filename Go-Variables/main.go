package main

import "fmt"

func main() {
	//Three ways to declare a variable in Go
	// 1. Using var keyword
	var num int = 1029

	fmt.Println("Num:", num)

	// 2. Using short variable declaration
	num2 := 1029
	fmt.Println("Short Variable Num2:", num2)

	// 3. Using var keyword with type inference
	var num3 = 1029
	fmt.Println("Variable Num3 with var keyword without type:", num3)
	otherThingAboutVariables()

	//constants variables and there is no need to use var keyword
	// and no need to use type and :=
	const num5 = 1029
	fmt.Println("Constant Num5:", num5)

}

//Camel case is used for variable names in Go
func otherThingAboutVariables() {
	//default values of variables
	var num4 int
	fmt.Println("Default value of num4:", num4) // 0

	//changing the value of a variable
	num4 = -99
	fmt.Println("Changed value of num4:", num4) // -99

	//multiple variable declaration like python
	var a, b, c = "Atul", 2, 3.5
	fmt.Println("Multiple variable declaration:", a, b, c) // 1 2 3
}

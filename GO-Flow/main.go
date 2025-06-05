package main

import "fmt"

func main() {
	var falseVal bool = false
	var trueVal bool = true

	fmt.Println("falseVal:", falseVal)
	fmt.Println("trueVal:", trueVal)

	fmt.Println("CompareBool(false, false):", CompareBool(1, 0))
	//switchTest()

	//switchTestMultiValue()
	//switchFallthrough()

	//switchWithExpression()
	switchExpressionWithStatment()
}

func CompareBool(a int, b int) bool {
	return a == b
}

func switchTest() {
	var a int = 1
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is neither 1 nor 2")
	}
}

func switchTestMultiValue() {
	var a int = 2
	switch a {
	case 1, 2:
		fmt.Println("a is 1 for value 1 or 2")
	case 4:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is neither 1 nor 2")
	}
}

func switchFallthrough() {
	var a int = 2
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
		fallthrough
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is neither 1 nor 2")
	}
}

func switchWithExpression() {
	var a int = 2 // if a conditions are met, it will print the corresponding message
	// if a is less than 1, it will print "a is less than 1"
	switch {
	case a < 1:
		fmt.Println("a is less than 1")
	case a < 3:
		fmt.Println("a is less than 3")
	default:
		fmt.Println("a is greater than or equal to 3")
	}
}

func switchExpressionWithStatment() {
	switch a := 2; a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
		fallthrough
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is neither 1 nor 2")
	}
}

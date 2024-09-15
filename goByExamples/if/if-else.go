package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 7 or 8 are even")
	}

	if 8%2 == 0 && 7%2 == 0 {
		fmt.Println("either 7 or 8 are even")
	} else {
		fmt.Println("condition failed")
	}

	// we can declare a variable in if but the variable declared will only be available in the scope of the conditional and brances

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

}

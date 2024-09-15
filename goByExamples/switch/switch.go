package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Println("write", i, "as ")
	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println(i, "is a boolean")

		case int:
			fmt.Println(i, "is a integer")

		default:
			fmt.Printf("we dont know the type %T\n", t)
		}
	}

	whatAmi(1)
	whatAmi(false)
	whatAmi("Hey")
}

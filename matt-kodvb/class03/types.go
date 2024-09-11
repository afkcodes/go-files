package main

import "fmt"

// var a int
// var (
// 	b = 2
// 	f = 2.01
// )

func main() {
	a := 1
	b := 2.4

	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)

	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)

	// a=b // not allowed directly
	a = int(b)
	fmt.Printf("a: %8T %[1]v\n", a)

	// we cannot convert a boolean to zero or zero to boolean

}

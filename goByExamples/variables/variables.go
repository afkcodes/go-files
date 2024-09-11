package main

import "fmt"

var name = "afkcodes"

// f := "declared and initialized"  cannot be declared here

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // multiple variable declaration
	fmt.Println(b, c)

	var d = true // boolean
	fmt.Println(d)

	var e int // zero-valued
	fmt.Println(e)

	f := "declared and initialized"
	fmt.Println(f)

	fmt.Println(name)
}

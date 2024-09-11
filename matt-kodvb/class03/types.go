package main

import (
	"fmt"
	"os"
)

// var a int
// var (
// 	b = 2
// 	f = 2.01
// )

func main() {
	// a := 1
	// b := 2.4

	// fmt.Printf("a: %T %v\n", a, a)
	// fmt.Printf("b: %T %v\n", b, b)

	// fmt.Printf("a: %8T %[1]v\n", a)
	// fmt.Printf("b: %8T %[1]v\n", b)

	// // a=b // not allowed directly
	// a = int(b)
	// fmt.Printf("a: %8T %[1]v\n", a)

	// we cannot convert a boolean to zero or zero to boolean

	var sum float64
	var n int
	for {
		var val float64
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}
		sum += val
		n++
	}
	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Println("The Average is", sum/float64(n))
}

// call with command Line
/*
 ➜  class03 git:(main) ✗ go run .
 ➜  class03 git:(main) ✗ cat nums.txt | go run .
 ➜  class03 git:(main) ✗ go run . < nums.txt
*/

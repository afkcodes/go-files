package main

import (
	"fmt"
	"os"
	"sayhello"
)

func main() {
	fmt.Println(sayhello.Say(os.Args[1:]))
}

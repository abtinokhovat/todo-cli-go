package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("help entity")
		return
	}
	entity := os.Args[1]

	if len(os.Args) <= 2 {
		fmt.Println("help operation")
		return
	}
	operation := os.Args[2]

	fmt.Println("entity:", entity, "operation:", operation)
}

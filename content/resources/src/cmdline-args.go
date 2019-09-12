package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	fmt.Println("With program name: ", argsWithProg)
	fmt.Println("Just args:", argsWithoutProg)
	fmt.Println("3rd Argument:", arg)
}
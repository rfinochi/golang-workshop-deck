package main

import "fmt"

func main() {
	// printTower
	const ROWS = 5
	for i := 0; i < ROWS; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println("")
	}
}
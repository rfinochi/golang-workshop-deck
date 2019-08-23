package main

import "fmt"

func main() {
	// simpleBreak
	for i := 1; i <= 10; i = i + 1 {
		if i == 8 {
			break
		}
		fmt.Print(i, " ")
	}
}
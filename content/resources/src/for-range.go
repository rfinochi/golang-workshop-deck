package main

import "fmt"

func main() {
	question := "Hello, how are you doing?"
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}

package main

import "fmt"

func main() {
	a := [...]string{"IRL", "IND", "US", "CAN"}
	b := a

	b[1] = "CHN"

	fmt.Println("Original:", a)
	fmt.Println("Copy    :", b)

	fmt.Println(len(a))
	fmt.Println(len(b))
}
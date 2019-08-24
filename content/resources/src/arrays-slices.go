package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	sa := a[1: 4]

	fmt.Println("Before:", a)
	sa[0] = 22

	fmt.Println("After:", a)
}
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	sum := 0
	for i, v := range a {
	    fmt.Println("Index:", i, " Value is:", v)
	    sum += v
	}
	fmt.Println("Sum:", sum)
}
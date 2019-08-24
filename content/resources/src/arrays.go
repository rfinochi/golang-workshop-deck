package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)

	a[0] = 1
	a[1] = 2
	a[2] = 3
	//a := [3]int{1, 2, 3}
	fmt.Println(a)
}
package main

import "fmt"

func main() {
	fmt.Println(Product(2, 2))
}

func Product(a int, b int) (prod int) {
    prod = a * b
    return
}
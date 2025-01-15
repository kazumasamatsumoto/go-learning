package main

import (
	"fmt"
	"go-basics/calculator"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is my first Go program.")
	fmt.Println("I'm learning Go language.")
	fmt.Println("Sum of 10 and 20:", calculator.Sum(10, 20)+calculator.Offset)
}

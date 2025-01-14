package main

import "fmt"

func main() {
	fmt.Println("test")
	sl := []int{1, 2, 3}
	if len(sl) > 0 {
		fmt.Println("unreachable code")
	}
}

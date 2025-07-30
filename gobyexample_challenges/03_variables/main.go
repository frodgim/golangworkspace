package main

import "fmt"

func main() {
	// TODO: Declare two variables, swap their values, and print the result
	var a int = 5
	b := 10
	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

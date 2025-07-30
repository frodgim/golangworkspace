package main

import "fmt"

func main() {
	// TODO: Check if a number is even or odd and print the result
	var sliceOfNumbers = []int{1, 2, 3, 4, 5}
	for _, num := range sliceOfNumbers {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}

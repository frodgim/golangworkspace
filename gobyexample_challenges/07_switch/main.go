package main

import "fmt"

func main() {
	// TODO: Print the word representation of a number from 1 to 3 using switch
	// let's grab this number from stdin
	fmt.Println("Enter a number between 1 and 3:")
	var number int
	fmt.Scanln(&number)
	for number < 1 || number > 3 {
		fmt.Println("Invalid input. Please enter a number between 1 and 3.")
		fmt.Scanln(&number)
	}

	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
}

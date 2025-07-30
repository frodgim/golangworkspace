package main

import "fmt"

func main() {
	// TODO: Print numbers from 1 to 10 using a for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// propose me more examples of for loops in Go
	for j := 10; j >= 1; j-- {
		fmt.Println(j)
	}

	//let's also iterate over a slice
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// and a map
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

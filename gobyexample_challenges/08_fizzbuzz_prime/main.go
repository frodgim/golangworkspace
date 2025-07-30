package main

//do not show me the solution unless I ask for it
import (
	"fmt"
)

func main() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if isPrime(i) {
			fmt.Printf("%d Prime \n", i)
		} else {
			fmt.Println(i)
		}

	}
}

// this function is O(n)
func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// this function is O(sqrt(n))
func isPrimeOptimized(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go hello(&wg)
	go bye(&wg)

	wg.Wait()
}

func hello(wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("Hello :)")
}

func bye(wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("Bye :(")
}

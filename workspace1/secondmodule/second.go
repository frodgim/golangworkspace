package main

import (
	// "fmt"
	"log"

	co "example/codeorgmodule/main2"
	"example/firstmodule"
)

func main() {
	log.SetPrefix("Second module: ")
	log.SetFlags(0)

	log.Println("Running the second module...")
	var name string = "Paco"

	log.Println("Running first.SaySomething() ...")

	first.SaySomething(name)

	log.Println("Running first.Hello() ...")
	// message, err := first.Hello("")
	message, err := first.Hello(name)

	// if err != nil {
	// 	fmt.Printf("We got an error: %s \n", err.Error())
	// }
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("fn 	: %s", message)

	log.Println("Calling Code Org Module...")
	log.Println(co.GetAnotherString())
}

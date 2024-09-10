package main

import (
	// "fmt"
	"example/codeorgmodule/main2"
	"log"
)

func main() {
	log.SetPrefix("Code Org:")
	log.SetFlags(0)
	log.Println("Executing Code Org main...")

	log.Println(GetString())
	log.Println(main2.GetAnotherString())
}

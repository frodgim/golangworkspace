package first

import (
	"fmt"
	"errors"
	"rsc.io/quote"
)

func main() {

	var phrase string = quote.Go()
	var secondPhrase string = quote.Opt()
    fmt.Println(phrase)
	fmt.Printf("quote.Opt() => %s", secondPhrase)
}

func SaySomething(name string) {
	fmt.Printf("Hi %s, I'm gonna say something probably with no sense \n => %s \n", name, quote.Go())
}

func Hello(name string) (string, error){
	if name == "" {
		return "", errors.New("any name was provided")
	}

	message := fmt.Sprintf("Hello %s", name)

	return message, nil

}
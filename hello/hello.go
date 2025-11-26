package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/ryderdain/gadgets/greetings"
)

func main() {
	message := gadgets.greetings.Hello("Moogle")
	fmt.Println(message)
	fmt.Println(quote.Go())
}

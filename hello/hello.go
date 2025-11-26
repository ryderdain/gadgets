package main

import (
	"fmt"
	"rsc.io/quote/v4"
	"github.com/ryderdain/gadgets"
)

func main() {
	message := gadgets.greetings.Hello("Moogle")
	fmt.Println(message)
	fmt.Println(quote.Go())
}

package main

import (
	"fmt"
	"rsc.io/quote/v4"
	"github.com/ryderdain/gadgets/pkg/greeting"
)

func main() {
	message := gadgets.Hello("Moogle")
	fmt.Println(message)
	fmt.Println(quote.Go())
}

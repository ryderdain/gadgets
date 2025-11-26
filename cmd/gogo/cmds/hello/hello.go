package hello

import (
	"fmt"
	"rsc.io/quote/v4"
	"github.com/ryderdain/gadgets/pkg/greetings"
)

func Hello() {
	message := greetings.Hello("Moogle")
	fmt.Println(message)
	fmt.Println(quote.Go())
}

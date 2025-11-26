package main

import "fmt"
import "rsc.io/quote/v4"

func main() {
    fmt.Println("Hello, kind stranger. Here's a quote for you:\n")
    fmt.Println(quote.Go())
}

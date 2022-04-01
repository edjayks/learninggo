package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Quote of the day: %v \n", quote.Go())
}

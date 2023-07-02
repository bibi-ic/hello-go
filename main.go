package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	msg := "Hi Golang !"
	fmt.Println(msg)

	msg2 := quote.Go()
	fmt.Println(msg2)
}

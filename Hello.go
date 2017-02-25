package main

import (
	"fmt"
)

func main() {
	message := "Hello Go world"

	greeting  := &message

	*greeting = "Hi"

	fmt.Println(message, *greeting)
}

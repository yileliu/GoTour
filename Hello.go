package main

import (
	"fmt"
)

func main() {
	 message := "Hello Go world"

	var greeting *string = &message

	fmt.Println(message, *greeting)
}

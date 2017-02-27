package main

import (
	"fmt"
)

type Salutation struct{
	name string
	greeting string
}

const (
	A = iota
	B
	C
)

func Greet(salutation Salutation) {
	fmt.Println(salutation.name)
	fmt.Println(salutation.greeting)
}

func main() {
	
	var s = Salutation{name: "Neal", greeting: "Hello"}

	Greet(s)
}



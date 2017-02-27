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
	_, alternate := CreateMessage(salutation.name, salutation.greeting);
	
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "HEY! " + name
}

func main() {
	
	var s = Salutation{name: "Neal", greeting: "Hello"}

	Greet(s)
}



package main

import (
	"fmt"
)

type Salutation struct{
	name string
	greeting string
}

type Printer func(string) ()

const (
	A = iota
	B
	C
)

func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo");
	
	do(message)
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "HEY! " + name
	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string){
		fmt.Println(s, custom)
	}
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func main() {
	
	var s = Salutation{name: "Neal", greeting: "Hello"}

	Greet(s, CreatePrintFunction("!!!"))
}



package greeting

import (
	"fmt"
)

type Salutation struct{
	Name string
	Greeting string
}

type Printer func(string) ()

const (
	A = iota
	B
	C
)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting);

	if prefix := "Mr "; isFormal{
		do(prefix + message)
	}
	do(alternate)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
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
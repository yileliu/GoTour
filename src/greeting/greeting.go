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

	if prefix := GetPrefix(salutation.Name); isFormal{
		do(prefix + message)
	} else{
		do(alternate)
	}
}

func GetPrefix(name string)(prefix string) {
	switch {
		case name == "Bob": 
			prefix = "Mr "
		case name == "Joe", name == "Amy", len(name) == 10: 
			prefix = "Dr "
		case name == "Mary": 
			prefix = "Miss "
		default: prefix = "Dude"
	}

	return
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
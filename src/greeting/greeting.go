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

func Greet(salutation Salutation, do Printer, isFormal bool, times int) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)

	i := 0
	for  i < times {
		if prefix := GetPrefix(salutation.Name); isFormal{
			do(prefix + message)
		} else{
			do(alternate)
		}

		i++ 
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

func TypeSwitchTest(x interface{}) {
	switch x.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case Salutation:
			fmt.Println("Salutation")
		default: 
			fmt.Println("Unknown")
	}
	
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
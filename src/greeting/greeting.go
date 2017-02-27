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

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range salutation{
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal{
			do(prefix + message)
		} else{
			do(alternate)
		}

	}
}

func GetPrefix(name string)(prefix string) {
	prefixMap := map[string]string{
		"Bob": "Mr ",
		"Joe": "Dr ",
		"Amy": "Dr ",
		"Mary": "Dr ",
	}
	
	prefixMap["Joe"] = "Jr "
	delete(prefixMap, "Mary")

	return prefixMap[name]
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
		fmt.Println(s + custom)
	}
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}
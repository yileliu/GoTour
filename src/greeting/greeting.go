package greeting

import (
	"fmt"
)


type Renamable interface{
	Rename(newName string)
}

type Salutation struct{
	Name string
	Greeting string
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func (salutation *Salutation) Write(p []byte)(n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

type Salutations []Salutation

type Printer func(string) ()

func (salutations Salutations) Greet( do Printer, isFormal bool, times int) {
	for _, s := range salutations {
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

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude"
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
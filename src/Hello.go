package main

import ("greeting")


func main() {
	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Amy", "What's up?"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("!!!"), true, 5)
	greeting.TypeSwitchTest("Bob")

}



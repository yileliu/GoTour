package main

import "./greeting"


func main() {
	
	var s = greeting.Salutation{"1234567890", "Hello"}

	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)
	greeting.TypeSwitchTest("Bob")
}



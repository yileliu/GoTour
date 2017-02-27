package main

import "./greeting"


func main() {
	
	var s = greeting.Salutation{"Amy", "Hello"}

	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), true)
}



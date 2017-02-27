package main

import "./greeting"


func main() {
	
	var s = greeting.Salutation{"Neal", "Hello"}

	greeting.Greet(s, greeting.CreatePrintFunction("!!!"), false)
}



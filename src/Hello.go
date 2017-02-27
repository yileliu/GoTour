package main

import ("greeting")


func main() {
	var s []int
	s = make([]int, 3)
	s[0] = 1


	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Amy", "What's up?"},
	}

	slice = append(slice, greeting.Salutation{"Frank", "Ho"})
	slice = append(slice[:1], slice[2:]...)


	//slice = slice[:2]

	greeting.Greet(slice, greeting.CreatePrintFunction("!!!"), true, 5)
	greeting.TypeSwitchTest("Bob")

}



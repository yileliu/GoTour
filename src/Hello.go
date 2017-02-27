package main

import ("greeting")


func main() {
	var s []int
	s = make([]int, 3)
	s[0] = 1


	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Amy", "What's up?"},
	}

	salutations = append(salutations, greeting.Salutation{"Frank", "Ho"})
	salutations = append(salutations[:1], salutations[2:]...)


	//slice = slice[:2]
	salutations[0].Rename("John")

	salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 5)
	greeting.TypeSwitchTest("Bob")

}



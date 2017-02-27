package main

import ("greeting"; "fmt")

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

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
	RenameToFrog(&salutations[1])
	fmt.Fprintf(&salutations[2], "The count is %d", 10)
	salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 5)
	greeting.TypeSwitchTest("Bob")

}



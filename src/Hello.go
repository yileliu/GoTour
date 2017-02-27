package main

import (
	"greeting"
	"fmt"
	"time"
)

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

	done := make(chan bool, 2)

	go func() {
		salutations.Greet(greeting.CreatePrintFunction("???"), true, 5)
		done <- true
		time.Sleep(100 * time.Millisecond)
		done <- true
		fmt.Println("Done !")
	}()

	salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 5)
	greeting.TypeSwitchTest("Bob")

	<- done
	
	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)
	go salutations.ChannelGreeter(c);
	go salutations.ChannelGreeter(c2);

	for {
		select {
			case s, ok := <- c: 
				if ok {
					fmt.Println(s, ":1")
				}else{
					return
				}
			case s, ok := <- c2: 
				if ok {
					fmt.Println(s, ":2")
				}else{
					return
				}
			

			default:
				fmt.Println("Waiting...")

		}
	}
}



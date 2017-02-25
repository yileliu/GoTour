package main

import (
	"fmt"
)

type Salutation struct{
	name string
	greeting string
}

func main() {
	
	var s = Salutation{name:"Neal", greeting:"Hello"}

	fmt.Println(s.greeting, ",", s.name)
}

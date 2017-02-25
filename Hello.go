package main

import (
	"fmt"
)

type Salutation struct{
	name string
	greeting string
}

const (
	A = iota
	B
	C
)

func main() {
	
	var s = Salutation{name:"Neal", greeting:"Hello"}
	s.name = "Y"

	fmt.Println(A, B, C)
}

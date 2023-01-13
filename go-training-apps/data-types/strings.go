package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings are mutable but only get declared once
	//below is decalaring a string with a value then changing it
	s1 := "Im a string"
	s1 = "I've been changed!"

	fmt.Println(s1)

	//strings have built in functions called like strings.function()
	//below is comparing two strings to see if 1 contains the other returning a boolean (case sesitive)
	c1 := "My name is Patrick"
	c2 := "Patrick"

	fmt.Println(strings.Contains(c1, c2))

	//replace all strings function, and split strings function examples
	r1 := "peter piper picked some peppers"

	fmt.Println(strings.ReplaceAll(r1, "p", "zz"))

	fmt.Println(strings.Split(r1, " ")) //this is now a list of strings

	//concat is done with simple plus sign
	fmt.Println(c1 + s1 + "annnnddd..." + r1)

}

package main

import "fmt"

func main() {
	fmt.Println("hello world")

	//if, else, for, switch, break
	// && (and) , || (or)

	//in web dev the pointers are used alot to ensure there is memory being allocated to things
	f := true
	flag := &f

	if flag == nil {
		fmt.Println("value is nil")
	} else if *flag {
		fmt.Println("true")
	} else if !*flag {
		fmt.Println("false")
	} else {
		fmt.Println("im not sure what you are")
	}

	//for loops (similar to C#)
	for y := 0; y < 10; y++ {
		fmt.Println(y)
	}

	//infinit loop
	// for {
	// 	fmt.Println("I will never stop")
	// }

	//loop over array
	a := []string{"patrick", "alex", "laguardia"}

	for x, v := range a {
		fmt.Println(x) // prints the index
		fmt.Println(v) // prints the value
	}

	//loop over a map of something we dont know yet
	mymap := make(map[string]interface{})
	mymap["first"] = "patrick"
	mymap["last"] = "laguardia"
	mymap["age"] = 32

	for x, y := range mymap {
		//fmt.Println(x, y)
		fmt.Printf("my key is: %s and my value is %v", x, y) //printing a formated string, always has to be %s and %v
	}

}

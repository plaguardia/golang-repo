package main

import "fmt"

func main() {
	//below is implicit data type assignment (assumes int)
	var (
		impnum  = 5
		impnum2 = 10
	)

	fmt.Println("implicit:", impnum+impnum2)

	//below is showing type casting to convert two diff data types to the same
	var cast1 int32
	var cast2 int64
	cast1 = 5
	cast2 = 10

	fmt.Println("typeCast:", cast1+int32(cast2))

	//below is most simple single line declaration of vars
	single1 := 5
	single2 := 10

	fmt.Println("single:", single1+single2)
}

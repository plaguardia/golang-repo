package main

import "fmt"

func swap(i1, i2 *int) {
	var temp int
	temp = *i2
	*i2 = *i1
	*i1 = temp
}

func main() {
	//pointers are vars that store the address in memory of another var
	m1 := 2
	ptr := &m1

	//prints hex of where value is stored in memory
	fmt.Println(ptr)

	// & (reference operator) * (dereference operator)
	//* prints the value behind the hex
	fmt.Println(*ptr)

	//use a swap function to switch the values behind the hex in memory
	i1, i2 := 500, 1000

	fmt.Println("PreSwap:", i1, i2)
	swap(&i1, &i2)
	fmt.Println("PostSwap:", i1, i2)
}

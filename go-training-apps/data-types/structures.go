package main

import "fmt"

// structure is an abstract data type that you can design with related types of data
// this is called "encapsulation"
type car struct {
	//below are refered to as member data types
	make  string
	model string
	year  int
}

// you can attach specific functions ("methods") to a data structure. "what are the things a 'car' can do"
func (x car) drive() {
	fmt.Println("driving...")
}

func (x car) stop() {
	fmt.Println("stopping...")
}

func (x car) print() {
	fmt.Println(x)
}

// in this example string is the expected data type of the value returned
func (x car) getMake() string {
	return x.make
}

func main() {
	//declare
	c1 := car{"subaru", "legacy", 2016}

	//object notation declaration
	c2 := car{
		make:  "chevy",
		model: "impreza",
		year:  2012,
	}

	fmt.Println(c1, c2)

	//calling your struct functions
	c1.drive()
	c1.stop()
	c2.print()
	fmt.Println(c2.getMake())

}

package main

import "fmt"

// interfaces define certain sets of functions that you want other structures to follow or have
// every different structure of a car better be able to drive and stop
// the value is huge bc you can define what specific structures are able to do. If you have a video game a charater needs to be able to move, attack, etc
type car interface {
	drive()
	stop()
}

// below "newcar" function is where the "magic" happens connecting/enforcing your interface
// below code would fail until we define BOTH a drive and a stop function for the struct of chevy based on the "car" interface being applied in the function
// we get this error/warning in code (until the stop method is created)

//cannot use (chevy literal) (value of type chevy) as car value in return statement: chevy does not implement car (method drive has pointer receiver)compilerInvalidIfaceAssign

func newcar(x string) car {
	return &chevy{x}
}

type chevy struct {
	chevymodel string
}

type subaru struct {
	subarumodel string
}

func (c *chevy) drive() {
	fmt.Println("chevy is on the move")
	fmt.Println(c.chevymodel)
}

func (c *chevy) stop() {
	fmt.Println("chevy is stopping")
}

func (s *subaru) drive() {
	fmt.Println("subaru is on the move")
	fmt.Println(s.subarumodel)
}

func main() {
	car1 := subaru{
		subarumodel: "impreza",
	}
	car2 := chevy{"impala"}

	car1.drive()
	car2.drive()
}

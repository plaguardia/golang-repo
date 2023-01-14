package main

import (
	"fmt"
	"time"
)

// function to wait 10 seconds
func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("heavy")
	}
}
func super_heavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("super heavy")
	}
}

func main() {

	// go routines are not "functions" the yare the way we call functions
	// in the ex below we are running our heavy function as a go routine, which allows it to run in the "background"  async with other things processed
	// that is why you see it finish much quicker
	go heavy() // running our heavy function as a GO Routine, notice the diff? SO MUCH FASTER!!
	go super_heavy()
	fmt.Println("Finished")
	time.Sleep(time.Second * 5)

	//select will be talked about more in channels section, but for now know that the below line will keep the main function alive indefinitly
	select {}

}

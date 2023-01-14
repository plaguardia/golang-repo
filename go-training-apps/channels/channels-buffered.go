package main

import "fmt"

type car struct {
	model string
}

func main() {
	//make() function kinda means "make" some space in memory for this
	//make can take another param that is "buffer size" (0 is included so this is actually 4)
	myChannel := make(chan int, 3)

	//send somethings on channel
	go func() {
		myChannel <- 1
		myChannel <- 2
		myChannel <- 3
		myChannel <- 4
		close(myChannel) // have to close the channel for for loop sniffer knows when to stop
	}()

	// you can roll over the range of a channel in a for look
	for x := range myChannel {
		fmt.Println(x)
	}

	myCarChan := make(chan *car, 3)

	go func() {
		myCarChan <- &car{"subaru"}
		myCarChan <- &car{"chevy"}
		myCarChan <- &car{"caddie"}
		close(myCarChan)
	}()

	for y := range myCarChan {
		fmt.Println(y.model)
	}

}

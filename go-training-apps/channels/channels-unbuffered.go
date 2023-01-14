package main

import "fmt"

func main() {
	//make() function kinda means "make" some space in memory for this
	//make a channel with nothing in in
	myChannel := make(chan int)
	fmt.Println(myChannel) //prints the memory allocated for this

	//send something on channel
	go func() {
		myChannel <- 1 // notice the arrow is pointing to the channel
	}()

	//sniff for something on channel
	val := <-myChannel
	fmt.Println(val)

	//send something on channel
	go func() {
		myChannel <- 2 // notice the arrow is pointing to the channel
	}()

	//sniff for something on channel
	val = <-myChannel
	fmt.Println(val)

}

package main

import (
	"fmt"
	"os"
	"time"
)

func waiting(stuff chan int, stop chan struct{}) {
	//select blocks are like switch cases but specifically for channels and are async (faster)
	// this func is an infinit loop until it gets told to stop
	for {
		time.Sleep(time.Second)
		select {
		case <-stuff:
			fmt.Println("there is stuff in the stuff queue")
		case <-stop:
			fmt.Println("someone told me me stop")
			os.Exit(0)
		}
	}
}

func main() {
	apples := make(chan int)
	noMoreApples := make(chan struct{})

	go func() {
		apples <- 1
		noMoreApples <- struct{}{}
	}()

	go waiting(apples, noMoreApples)

	//now you should understand why empty set in select block keeps main running indef
	select {}

}

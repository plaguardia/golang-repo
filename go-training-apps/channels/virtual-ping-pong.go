package main

import (
	"fmt"
	"time"
)

// The player1 prints a ping and waits for a pong
func player1(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 0
	}
}

// The player2 prints a pong and waits for a ping
func player2(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 0
	}

}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go player1(ping, pong)
	go player2(ping, pong)

	//start the game by sending a number into the ping channel (then they start to send it back and forth)
	ping <- 0

	select {} // run indeff

}

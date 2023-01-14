package main

import (
	"fmt"
	"sync"
)

func main() {

	//wait groups these give functions that have been added to the group the ability to wait or be done

	wg := &sync.WaitGroup{}
	wg.Add(2) // this being 2 means that it will wait to see 2 wg.done() calls before moving on
	// the way these functions are called is refered to as lambda function or anonomous functions (use once and throw away or discard)
	go func() {
		fmt.Println("hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("world")
		wg.Done()
	}()
	fmt.Println("start") // bc this is before the wg.wait() it will not wait for it
	wg.Wait()            // this says to wait for the wait group to be done before moving on
	fmt.Println("finished")

}

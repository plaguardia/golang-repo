package main

import "fmt"

func main() {
	//fmt.Println("hello world")

	for i := 0; i < 10; i++ {
		if i%2 == 0 { //%2 means divisible by 2 - this returns a 0 for true and a 1 for false
			continue //continue inside a for loop means to start the loop from the top without doing the rest of the loop
		}
		fmt.Println(i)
	}

	flag := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 { //%2 means divisible by 2 - this returns a 0 for true and a 1 for false
			flag = false
			break //break inside a for loop means to completely stop processing the loop and break out of it
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	//switch style 1
	day := "fri"
	switch day {
	case "fri":
		fmt.Println("tgif")
	case "mon", "tues", "wed":
		fmt.Print("boring")
	default:
		fmt.Println("default")
	}

	//switch style 2
	switch {
	case day == "fri":
		fmt.Println("tgif")
		break
	}

}

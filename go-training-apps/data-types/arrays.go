package main

import "fmt"

func todo() {
	//declaring array
	arr := []int32{1, 2, 3}
	arr2 := []string{"doe", "ray", "me"}

	//appending to array
	arr2 = append(arr2, "fa", "so")

	fmt.Println(arr, arr2)
}

func main() {
	todo() //notice we are calling a separate function instead of coding in our main func
}

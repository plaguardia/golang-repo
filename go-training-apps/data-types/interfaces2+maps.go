package main

import "fmt"

// empty interface is a "black box" it can be anything of any type
// so if you don't know the type of the arg in your function during compile time you can set it to a blank interface
func anything(x interface{}) {
	fmt.Println(x)
}

func main() {
	//fmt.Println()
	anything(3.14)
	anything("its meeee mario")
	anything(struct{}{}) //first {} declares the struct, second {} instatiates it to empty. Non empty would look like struct{person string}{"kevin"}

	//declare a map (key value pairs) with a string key and an interface value bc we don't know what type the value is going to be
	m := make(map[string]interface{})
	m["name"] = "Patrick"
	m["age"] = 32

	fmt.Println(m)

}

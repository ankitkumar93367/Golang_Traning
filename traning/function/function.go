package main

import "fmt"

/*
func main() {
	a := func() {
		fmt.Println("heloo world")
	}
	a()
	fmt.Printf("%T", a)
}
*/

func main() {
	func(a string) {
		fmt.Println(a)
	}("ankit")
}

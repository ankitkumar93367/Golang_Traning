package main

import (
	"fmt"
)

func addValue(a int, b int, additionChannel chan int) {

	c := a + b
	additionChannel <- c
}

func multiValue(a int, b int, multiChannel chan int) {

	c := a * b
	multiChannel <- c
}

func main() {

	fmt.Println("enter the value A")
	var a int
	fmt.Scan(&a)

	fmt.Println("enter the value B")
	var b int
	fmt.Scan(&b)

	addition := make(chan int)
	multiply := make(chan int)

	go addValue(a, b, addition)
	go multiValue(a, b, multiply)

	add, multi := <-addition, <-multiply

	fmt.Println("Final output", add, multi)

}

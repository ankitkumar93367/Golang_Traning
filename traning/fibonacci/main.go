package main

import (
	"fmt"
)



func fibonacci(ch chan int, quit chan bool) {
    x, y := 0, 1
    for {
        select {
        case ch <- x: // write to channel ch
		fmt.Println("fibonacci genrate")
            x, y = y, x+y
			sent <- true
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}
var sent = make(chan bool)
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func (n int)  {
		for i :=0; i <= 10; i++{
			fmt.Println(<-ch) //reading the channel
			<-sent
		}
		quit <- false // send signal to fibonacci
	}(n)

	fibonacci(ch, quit)
}
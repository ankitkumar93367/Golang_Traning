package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	fmt.Println("Sending value to channel")
	go send(ch)

	go reveiver(ch)
	time.Sleep(time.Second * 1)
}

func send(ch chan int) {
	fmt.Println("send")
	ch <- 12
}

func reveiver(ch chan int) {
	fmt.Println("reveiver")
	val := <-ch
	fmt.Printf("Value Received=%d in receive function\n", val)

}

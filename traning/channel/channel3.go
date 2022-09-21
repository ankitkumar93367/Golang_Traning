package main

import (
	"fmt"
)

func producer(chal chan int) {

	for i := 0; i < 10; i++ {
		chal <- i
	}

	close(chal)

}

func main() {
	ch := make(chan int)

	go producer(ch)

	res := <-ch

	fmt.Println("response", res)

	for v := range ch {
		fmt.Println(v)
	}

	/*	for {
		v, ok := <-ch

		if ok == false {
			break
		}
		fmt.Println("received", v, ok)
	}*/

}

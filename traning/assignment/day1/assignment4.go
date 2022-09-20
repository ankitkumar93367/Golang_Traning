package main

import "strings"

/*
Generate the Series 1A2B3C4D5E....... using  goroutine and channel  .Routine 1 shall generate A,B C D and routine 2 shall
generate 1 ,2,3,4,5,6....... Both this output join together in the output
*/



func main () {

	number := []int{1, 2, 3, 4, 5}
	alphabets := []string{ "A","B", "C","D", "E" }


	myChannel := make(chan string)


}
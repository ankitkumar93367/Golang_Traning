package main

import "fmt"

/*
Given an integer n, return an array ans of length n + 1 such that for each i(0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

*/

func main() {

	fmt.Println("enter the number...")
	var number int
	fmt.Scan(&number)

	c := make([]string, number)

	for i := 0; i < number; i++ {
		c[i] = fmt.Sprintf("%b", i)
	}

	fmt.Println(c)

}

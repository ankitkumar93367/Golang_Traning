package main

import (
	"fmt"
)

//Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the
func main() {

	num1 := 10
	num2 := 5

	for num2 > 0 {
		num1++
		num2--
	}

	fmt.Println(num1)

}

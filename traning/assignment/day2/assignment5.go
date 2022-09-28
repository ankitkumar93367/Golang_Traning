package main

import "fmt"

func LargestNumber(arr []int) int {

	var largerNumber, temp int
	for _, element := range arr {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	return largerNumber
}
func main() {

	number := LargestNumber([]int{1, 3, 2, 5, 4, 6, 9, 8})

	fmt.Println("large number", number)
}

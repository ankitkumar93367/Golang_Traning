package main

import (
	"fmt"
)

/*
Input: n = 2
Output: [0,1,1]
*/

func main() {

	ans := make([]int, 2)

	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			ans[i] = ans[i/2]
		} else {
			ans[i] = ans[i/2] + 1
		}
	}

	fmt.Println(ans)
}

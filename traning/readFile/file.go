package main

import (
	"fmt"
	"os"
)

func main() {
	contents, err := os.ReadFile("sample3.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(contents))
	}

}

package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID int
	Name string
	Password string
	LoggedAt time.Time
}

func main() {

	user := User{}
	user.ID = 1232
	user.Name = "peter"
	user.Password = "1234"
	user.LoggedAt = time.Now()

	content, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("userFile.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
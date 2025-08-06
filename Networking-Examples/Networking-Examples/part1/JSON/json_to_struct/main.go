package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Title      string `json:"title"`
	Instructor string `json:"author"`
}

func main() {
	jsonClass := `{"title":"Golang Introduction","author":"Travis"}`

	var itmd469 Class

	err := json.Unmarshal([]byte(jsonClass), &itmd469)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(itmd469.Instructor)
	fmt.Println(itmd469.Title)
}

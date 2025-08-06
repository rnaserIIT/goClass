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
	itmd469 := Class{Title: "Golang Introduction", Instructor: "Travis"}

	jsonClass, err := json.Marshal(itmd469)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonClass))

}

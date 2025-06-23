package main

import (
	"fmt"
	"strings"
)

func main() {
	var live bool
	live = true
	var response string
	for live {
		println("how are you doing today?")
		println("For a list of commands and options type !commands or to exit type kill")
		fmt.Scanln(&response)
		strings.ToLower(response)
		switch response {
		case "good":
			fmt.Println("great to hear!")

		case "bad":
			fmt.Println("I'm sorry to hear that, maybe you should watch a funny movie or go for a walk to cheer yourself up.")
		case "tired":
			fmt.Println("Try to get some rest soon—you deserve it!")
		case "average":
			fmt.Println("Hope things improve from here.")
		case "!commands":
			fmt.Println(" good, bad, tired, average,!commands, kill")
		case "kill":
			live = false
			fmt.Println("goodbye")
		default:
			fmt.Println("invalid option")

		}

	}

}

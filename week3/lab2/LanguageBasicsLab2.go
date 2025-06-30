package main

import (
	"fmt"
	"lab2/mysum"
)

func main() {
	var num1 int
	var num2 int
	println("please give me 2 integers seperated by a space to add together and I'll tell you if the sum is even or odd")
	fmt.Scanln(&num1, &num2)
	println(mysum.Sum(num1, num2))
	println(num1, num2)

}

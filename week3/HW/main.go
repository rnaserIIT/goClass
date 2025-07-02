package main

import (
	"HW/MyMath"
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var operator string
	//var answer float64

	fmt.Println("please give me the first number youd like to calculate")
	fmt.Scanln(&num1)
	fmt.Println("please give me the second number youd like to calculate")
	fmt.Scanln(&num2)
	fmt.Println("press 1 to add 2 to subtract 3 to multiply and 4 to divide")
	fmt.Scanln(&operator)
	switch operator {
	case "1":
		fmt.Println(MyMath.Add(num1, num2))

	case "2":
		fmt.Println(MyMath.Sub(num1, num2))
	case "3":
		fmt.Println(MyMath.Mult(num1, num2))
	case "4":
		fmt.Println(MyMath.Div(num1, num2))
	default:
		fmt.Println("error please run the program again")

	}

}

/*
func add(a, b float64) float64 {
	return (a + b)

}

func sub(a, b float64) float64 {
	return (a - b)

}

func mult(a, b float64) float64 {
	return (a * b)

}

func div(a, b float64) float64 {
	return (a / b)

}

*/

package main

import "fmt"

type Employee struct {
	name          string
	Age           int
	position      string
	MonthlySalary float64
}

func (E Employee) CalcYearlyCost() float64 {
	return (E.MonthlySalary * 12)
}

func (E Employee) PrintInfo() {
	fmt.Println("Name:", E.name)
	fmt.Println("Age:", E.Age)
	fmt.Println("Position:", E.position)
	fmt.Println("MonthlySalary:", E.MonthlySalary)

}
func main() {
	boss := Employee{}
	boss.name = "Kevin"
	boss.Age = 53
	boss.position = "head of IT"
	boss.MonthlySalary = 16000
	boss.PrintInfo()
	println("this Employee's yearly salary is ", boss.CalcYearlyCost())

}

package main

import "fmt"

type Government struct {
	CorpTax       float64
	TotalEarnings float64
	TotalCost     float64
}

func (G Government) CalcYearlyCost() float64 {
	answer := (G.TotalEarnings - G.TotalCost) * G.CorpTax
	return (answer)
}

type Employee struct {
	name          string
	Age           int
	position      string
	MonthlySalary float64
	TotalCost     float64
}

func (E Employee) CalcYearlyCost() float64 {
	return (E.MonthlySalary * 12)
}

type Facilities struct {
	Type          string
	RentCost      float64
	MaitnenceCost float64
	TotalCost     float64
}

func (F Facilities) CalcYearlyCost() float64 {
	return ((F.RentCost * 12) + F.MaitnenceCost)
}

type Cost interface {
	CalcYearlyCost() float64
}
type TotalCosts struct {
	Facilities
	Employee
	Government
}

func CalcYearlyCost(c Cost) float64 {
	return c.CalcYearlyCost()

}

func main() {
	boss := Employee{}
	boss.name = "Kevin"
	boss.Age = 53
	boss.position = "head of IT"
	boss.MonthlySalary = 16000

	fmt.Println("this Employee's yearly cost is ", CalcYearlyCost(boss))
	boss.TotalCost = CalcYearlyCost(boss)

	warehouse := Facilities{Type: "Storage", RentCost: 6500.25, MaitnenceCost: 1000}

	fmt.Println("the warehouse facility is for ", warehouse.Type)
	fmt.Println("the yearly cost of the facility is ", CalcYearlyCost(warehouse))
	warehouse.TotalCost = CalcYearlyCost(warehouse)
	CostOfBuis := (CalcYearlyCost(boss) + CalcYearlyCost(warehouse))
	fmt.Println("the cost of the emeployee and the facility is ", CostOfBuis)

	pig := Government{CorpTax: .21, TotalEarnings: 999999.55, TotalCost: CostOfBuis}

	EVERYTHING := TotalCosts{
		Employee: Employee{
			name: "chuck", Age: 35, position: "helpdesk", MonthlySalary: 4000,
		},
		Facilities: Facilities{
			Type: warehouse.Type, RentCost: warehouse.RentCost, MaitnenceCost: warehouse.MaitnenceCost,
		},
		Government: Government{
			CorpTax: pig.CorpTax, TotalEarnings: pig.TotalEarnings, TotalCost: pig.TotalCost,
		},
	}
	profit := EVERYTHING.Government.TotalEarnings -
		(EVERYTHING.Employee.TotalCost) -
		(EVERYTHING.Facilities.TotalCost) -
		(EVERYTHING.Government.TotalCost)

	fmt.Println("total profit: $", profit)

}

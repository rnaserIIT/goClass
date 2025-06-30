package main

type Employee struct {
	name          string
	Age           int
	position      string
	MonthlySalary float64
}

func (E Employee) CalcYearlyCost() float64 {
	return (E.MonthlySalary * 12)
}

type Facilities struct {
	Type          string
	RentCost      float64
	MaitnenceCost float64
}

func (F Facilities) CalcYearlyCost() float64 {
	return ((F.RentCost * 12) + F.MaitnenceCost)
}

type Cost interface {
	CalcYearlyCost() float64
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

	println("this Employee's yearly cost is ", CalcYearlyCost(boss))

	warehouse := Facilities{Type: "Storage", RentCost: 6500.25, MaitnenceCost: 1000}

	println("the warehouse facility is for ", warehouse.Type)
	println("the yearly cost of the facility is ", CalcYearlyCost(warehouse))

	println("the cost of the emeployee and the facility is ", (CalcYearlyCost(boss) + CalcYearlyCost(warehouse)))
}

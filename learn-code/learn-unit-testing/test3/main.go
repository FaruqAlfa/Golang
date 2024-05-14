package main

type Person struct {
    Name string
    Age int
    BaseSalary float64
    Wealth float64
}

func NewPerson(name string, age int, baseSalary float64) *Person {
    return &Person{
        Name: name,
        Age: age,
        BaseSalary: baseSalary,
    }
}

func (p *Person) GetWealth() float64 {
    return p.Wealth
}

func (p *Person) SetBonus(monthOfWork int) {
    p.Wealth += (float64(monthOfWork) / 12) * p.BaseSalary * 0.5
}

func (p *Person) AddIncome(income float64) {
    p.Wealth += income
}

func (p *Person) AddCost(cost float64) {
    p.Wealth -= cost
}

func (p *Person) ClearWealth() {
    p.Wealth = 0
}
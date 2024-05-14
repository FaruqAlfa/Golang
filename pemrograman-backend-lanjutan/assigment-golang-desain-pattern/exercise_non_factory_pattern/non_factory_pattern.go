package main

import "fmt"

type Element interface{
	Attribute() string
}

type ImaginaryElement struct{
	power int
}

func (e *ImaginaryElement) IncreasePower() {
	e.power++
}

func (e *ImaginaryElement) Attribute() string {
	return "Imaginary Element Attribute with power: " + fmt.Sprint(e.power)
}

type QuantumElement struct{
	power int
}

func (e *QuantumElement) IncreasePower() {
	e.power += 2
}
func (e *QuantumElement) Attribute() string {
	return "Quantum Element Attribute with power: " + fmt.Sprint(e.power)
}

func Attack(element Element) {
	fmt.Println(element.Attribute() + " Attack !!!")
}

func main() {
	ImaginaryElement := &ImaginaryElement{}
	ImaginaryElement.IncreasePower()
	Attack(ImaginaryElement)

	
	QuantumElement := &QuantumElement{}
	QuantumElement.IncreasePower()
	Attack(QuantumElement)

}
package main

import "fmt"

type Element interface {
	Attribute(power int) string
}

type ImaginaryElement struct{}

func (e *ImaginaryElement) Attribute(power int) string {
	return fmt.Sprintf( "Imaginary Element Attribute with power: %d", power)
}

type QuantumElement struct{}

func (e *QuantumElement) Attribute(power int) string {
	return fmt.Sprintf( "Quantum Element Attribute with power: %d", power)
}

type AttackDamage interface{
	DamageElement() (Element, int)
}

type ImaginaryAttackDamage struct{
	power int
}

func (iad *ImaginaryAttackDamage) IncreasePower() {
	iad.power++
}

func (iad *ImaginaryAttackDamage) DamageElement() (Element, int) {
	return &ImaginaryElement{}, iad.power
}

type QuantumAttackDamage struct{
	power int
}

func (qad *QuantumAttackDamage) IncreasePower() {
	qad.power += 2
}

func (qad *QuantumAttackDamage) DamageElement() (Element, int) {
	return &QuantumElement{}, qad.power
}

func ElementDamage(attackDamage AttackDamage) {
	element, power := attackDamage.DamageElement()
	fmt.Println(element.Attribute(power))
}

func main() {

	ImaginaryAttackDamage := &ImaginaryAttackDamage{}
	ImaginaryAttackDamage.IncreasePower()
	ElementDamage(ImaginaryAttackDamage)

	QuantumAttackDamage := &QuantumAttackDamage{}
	QuantumAttackDamage.IncreasePower()
	ElementDamage(QuantumAttackDamage)
}
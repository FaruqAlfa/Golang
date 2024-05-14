package main

import (
	"testing"
)

func TestElementDamage(t *testing.T) {
	imaginaryDamage := &ImaginaryAttackDamage{}
	imaginaryDamage.IncreasePower()
	result := testElementDamage(imaginaryDamage)
	expected := "Imaginary Element Attribute with power: 1"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	quantumDamage := &QuantumAttackDamage{}
	quantumDamage.IncreasePower()
	result = testElementDamage(quantumDamage)
	expected = "Quantum Element Attribute with power: 2"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func testElementDamage(attackDamage AttackDamage) string {
	element, power := attackDamage.DamageElement()
	return element.Attribute(power)
}

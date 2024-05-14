package main

import "testing"

func TestAttack(t *testing.T) {
	imaginaryElement := &ImaginaryElement{}
	imaginaryElement.IncreasePower()
	if imaginaryElement.Attribute() != "Imaginary Element Attribute with power: 1" {
		t.Errorf("Expected Imaginary Element Attribute with power: 1, but got %s", imaginaryElement.Attribute())
	}

	quantumElement := &QuantumElement{}
	quantumElement.IncreasePower()
	if quantumElement.Attribute() != "Quantum Element Attribute with power: 2" {
		t.Errorf("Expected Quantum Element Attribute with power: 2, but got %s", quantumElement.Attribute())
	}
}

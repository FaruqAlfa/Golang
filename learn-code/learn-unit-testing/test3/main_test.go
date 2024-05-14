package main

import "testing"

var person *Person = NewPerson("John", 20, 100000)

func SetupTest(t *testing.T) func(*testing.T) {
    t.Log("Setup test case")

    person.SetBonus(12)

    return func(t *testing.T) {
        t.Log("Clear data test case set wealth to 0")
        person.ClearWealth()
    }
}

func TestAddIncome(t *testing.T) {
    // menampung fungsi 'SetupTest' di variable 'clear'
    clear := SetupTest(t)

    // melakukan clear setelah menjalankan test case
    // kita dapat memastikan bahwa clear ini dijalankan setelah test case menggunakan 'defer'
    defer clear(t)

    person.AddIncome(100000)
    if person.GetWealth() != 150000.0 {
        t.Errorf("Expected %f, but got %f", 150000.0, person.GetWealth())
    }
}

func TestAddCost(t *testing.T) {
    clear := SetupTest(t)

    defer clear(t)

    person.AddCost(10000)
    if person.GetWealth() != 40000.0 {
        t.Errorf("Expected %f, but got %f", 40000.0, person.GetWealth())
    }
}
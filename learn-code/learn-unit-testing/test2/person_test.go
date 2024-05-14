// file person_test.go
package entity

import (
    "testing"
    "time"
)

func TestGetAge(t *testing.T) {
    var person = Person{
        Name: "John Doe",
        DateBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
    }

    var expected int = 44

    if person.GetAge() != expected {
        t.Errorf("Expected %d, but got %d", expected, person.GetAge())
    }
}
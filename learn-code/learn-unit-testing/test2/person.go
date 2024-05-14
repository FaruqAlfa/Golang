// file person.go
package entity

import "time"

type Person struct {
    Name string
    DateBirth time.Time
}

func (p Person) GetAge() int {
    return time.Now().Year() - p.DateBirth.Year()
}
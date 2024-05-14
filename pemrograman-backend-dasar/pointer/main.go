package main

import (
	"fmt"
)

func AddByOne(number *int) { // hanya menerima parameter pointer dari tipe data 'int'
    *number = *number + 1

	// number = number + 1

	fmt.Println("in function AddByOne:", number)
}

func main(){
	//Syntax &
	// var number = 10
    // fmt.Println(&number)

	//Syntax *
	// var number = 10
    // var pointerNumber = &number // pointerNumber berisi alamat memori dari variable 'number'

    // fmt.Println(pointerNumber)
    // fmt.Println(*pointerNumber) // mendapatkan nilai dari variable 'number'

	// var number = 10
    // var pointerNumber = &number // pointerNumber berisi alamat memori dari variable 'number'

    // fmt.Println(pointerNumber)
    // fmt.Println(*pointerNumber) // mendapatkan nilai dari variable 'number'

	// var number = 10
    // AddByOne(&number)
    // fmt.Println(number)

	//Pas by Value
	// var number int = 10

    // AddByOne(number)

    // fmt.Println("in function main:", number)

	//Pas by Reference
	// var number int = 10

    // AddByOne(&number)

    // fmt.Println("in function main:", number)

	//Point Zero Value
	var number int = 10

    AddByOne(&number)

    fmt.Println("in function main:", number)
}
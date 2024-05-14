package main

import "fmt"

func main() {
	// this is comment
	// print output "Hello World!"
	// fmt.Println("Hello World!");

	/*
	   this is comment
	   print output "Hello World!"
	   fmt.Println("Hello World!");
	*/
	//fmt.Println("Hello Go!");

	/*fmt.Println("This is output Println")
	    //fmt.Println("Hello Go!")
		fmt.Println("Hello", 1234, true)

		fmt.Print("This is output Print ")
	    fmt.Print("Hello Go!")

		fmt.Printf("My name is %s\n", "Budi") // %s is formatting for string data type
	    fmt.Printf("%d is my contact number\n", 123434)  // %d is formatting for integer / numberic data type*/

	//BASIC INPUT ARGUMENT

	/*var name, address string
	    fmt.Print("Enter your name and address : ")
	    fmt.Scan(&name, &address) // set input to variable 'name' and 'address'
		//Jika tidak di beri spasi maka data yang di inputkan akan di masukkan di variabel pertamanya
	    fmt.Println("Hello", name)
	    fmt.Println("Your address", address)*/

	var name, address string
	/*fmt.Print("Enter your name : ")
	  fmt.Scanln(&name)

	  fmt.Print("Enter your address : ")
	  fmt.Scanln(&address)

	  fmt.Println("Hello", name)
	  fmt.Println("Your address", address)*/

	fmt.Print("Enter your name and address : ")
	fmt.Scanf("%s %s", &name, &address) // menerima 2 input string yang dipisahkan oleh spasi, yang pertama akan diassign ke variable name dan yang kedua akan diassign ke variable address

	fmt.Println("Hello", name)
	fmt.Println("Your address", address)

}

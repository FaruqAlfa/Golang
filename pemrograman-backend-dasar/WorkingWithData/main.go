package main

import (
	"fmt"
	"strings"
	// "time"
	//"math"
	// "reflect"
	// "strconv"
)

func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func sayHello(name string) {
	fmt.Printf("Hello, %s", name)
}

func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func myHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

// BASIC DATA TYPE
func main() {
	//1. String
	//len digunakan untuk menghitung jumlah karakter string
	//fmt.Println(len("Hello World"))

	// fmt.Println("Golang"[0])
	// fmt.Println("Golang"[2])
	//Hasilnya akan berupa data byte yang berisi ASCII code yang mewakili karakter tersebut.

	//Jika ingin menampilkannya kembali dalam bentuk string, kita bisa menggunakan fungsi string dengan parameter yang diisi byte data.
	// fmt.Println(string("Golang"[0]))

	//2. NUMERIC
	//int
	// fmt.Println(-30)
	// fmt.Println(10_000_000_000)

	//float
	// fmt.Println(3.14)
	// fmt.Println(-2.232321431412)

	//OPERATION
	// fmt.Println(1 + 1)
	// fmt.Println(2 - 1)
	// fmt.Println(2 * 2)
	// fmt.Println(4 / 2)
	// fmt.Println(10 % 3) // sisa pembagian 10 dengan 3 adalah 1

	//AUGMENTED ASSIGNMENT
	// var a = 10
	// a += 10 // a = 10 + 10 = 20
	// fmt.Println(a)

	//UNARY OPERATOR
	// var a = 10
	// a++ // a = 10 + 1 = 11

	// fmt.Println(a)

	// var b = 10
	// fmt.Println(-b)

	//BOOLEAN
	// fmt.Println("Benar =", true)
	// fmt.Println("Salah =", false)

	///VARIABLE
	// hello := "Hello World"
	// println(hello)

	//CASTING / CONVERT
	// var number int32 = 10

	// var bigNum = int64(number) // convert int32 to int64

	// var floatNum = float32(number) // convert int32 to float32

	// fmt.Println(bigNum)
	// fmt.Println(floatNum)

	//fmt.Sprintf
	// var number int32 = 10
	// var isMaried bool = false

	// var numString = fmt.Sprintf("%d", number)        // convert int32 ke string
	// var isMariedString = fmt.Sprintf("%t", isMaried) // convert bool ke string
	// var Pi = fmt.Sprintf("%f", math.Pi)              // convert float64 dari variable 'Pi' di package 'math' ke string

	// fmt.Println(numString)
	// fmt.Println(isMariedString)
	// fmt.Println(Pi)

	// fmt.Println("type numString:", reflect.TypeOf(numString))
	// fmt.Println("type isMarriedString:", reflect.TypeOf(isMariedString))
	// fmt.Println("type Pi:", reflect.TypeOf(Pi))

	//strconv.itoa
	// var str = strconv.Itoa(10) // convert string to int

	// fmt.Println(str)
	// fmt.Println("type:", reflect.TypeOf(str))

	//strconv.Atoi
	// var str = strconv.Itoa(10) // convert string to int

	// fmt.Println(str)
	// fmt.Println("type:", reflect.TypeOf(str))

	//CONDITION
	// var name1, name2 = "John", "John"

	// var result = name1 == name2
	// var result2 = 10 < 2

	// fmt.Println(result)
	// fmt.Println(result2)

	// var score, attendance = 90, 70

	// var graduated = score > 80 && attendance > 80

	// fmt.Println(graduated)

	// if x, y := 75, 80; x > 70 && y > 70 {
	// 	fmt.Println("Selamat, anda lulus")

	// 	fmt.Println("Nilai rata-rata:", (x+y)/2)
	// }

	//SWITCH
	// day := 4
	// switch day {
	//     case 1, 2, 3, 4, 5:
	//         fmt.Println("Weekday")
	//     case 6, 7:
	//         fmt.Println("Weekend")
	//     default:
	//         fmt.Println("Invalid day")
	// }

	// number := 6
	// switch {
	//     case number > 10:
	//         fmt.Println("High number")
	//     case number > 5:
	//         fmt.Println("Middle number")
	//     case number > 1:
	//         fmt.Println("Low number")
	//     default:
	//         fmt.Println("Zero")
	// }

	//LOOPING
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Hello Go!", i+1)
	// }

	// //Break
	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Hello Go!", i+1)
	// }

	// //Continue
	// for i := 1; i <= 5; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}

	// 	fmt.Println("Hello Go!", i)
	// }

	// //Condition-only in for
	// i := 0 // init statement
	// for i < 5 {
	// 	fmt.Println("Hello Go!", i+1)
	// 	i++ // post statement
	// }

	//empty for statement
	// var i = 0 // init statement

	// for {
	// 	if i >= 5 { // condition statement
	// 		break // menghentikan looping jika kondisi 'true'
	// 	}

	// 	fmt.Println("Hello Go!", i+1) // code execute

	// 	i++ // post statement
	// }

	//Infinite loop
	// for {
	// 	fmt.Println("Hello Go!")

	// 	time.Sleep(1 * time.Second) // ini adalah fungsi yang dibuat agar program delay selama 1 detik
	// }

	//Nested looping
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("Hello Go! i: %d, j: %d\n", i, j)
	// 	}
	// }

	//FUNCTION
	result := addNumbers(21, 13)
	fmt.Println("Sum:", result)

	sayHello("Fahmi")

	fmt.Println(Sum(1, 2))                          // 3
	fmt.Println(Sum(1, 2, 3, 4, 5))                 // 15
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55

	myHobbies("Aditira", "Coding", "Gaming", "Jogging")
}

package main

import (
	"fmt"
)

// func apply(f func(x int) int, number int) int {
//     return f(number)
// }

//Closure
// func NewCounter() func() int {
//     i := 0 // nilai variabel `i` adalah closure bagi inner function `inc`
//     inc := func() int {
//         i++
//         return i
//     }
//     return inc
// }

//Factorial of a number using Go Recursion
func factorial (num int) int {
	if num == 0 {
	  return 1
	} else {
	  return num * factorial(num - 1)
	 }
  }

//Rekursive
// func countDown(number int) {
// 	// // menampilkan parameter `number`
// 	// fmt.Println(number)
// 	// // pemanggilan recursive dengan mengurangi nilai `number`
// 	// countDown(number - 1)

// 	if number > 0 {
// 		fmt.Println(number)
// 		countDown(number - 1)
// 	  } else {
// 		fmt.Println("Countdown Stops")
// 	  }
//   }


//Difference between Recursion and Iteration
func factorialRecursive(value int) int {
	if value == 1 {
	  return 1
	} else {
	  return value * factorialRecursive(value-1)
	}
  }
  
  func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
	  result *= i // result = result * i
	}
	return result
  }
  
func main() {

	//Function as Variable
	// Anonymous function assigned to variable inc:
	// inc := func(number int) int {
	// 	return number + 1
	// }

	// // Assigned variable function to variable calculate:
	// calculate := inc(5)
	// fmt.Println(calculate) // 6

		// Function as parameter
		// fungsi sebagai variable bernama `inc`
		// inc := func(number int) int {
		// 	return number + 1
		// }
		// // fungsi sebagai variable bernama `dec`
		// dec := func(number int) int {
		// 	return number - 1
		// }
	
		// fmt.Println(apply(inc, 5)) // 6
		// fmt.Println(apply(dec, 5)) // 4

		//CLOSURE
	// ctr := NewCounter()

    // fmt.Println(ctr()) // 1 => nilai variabel `i` sekarang 0 dan menghasilkan 1
    // fmt.Println(ctr()) // 2 => nilai variabel `i` sekarang 1 dan menghasilkan 2

	//IIFE (Immediately-invoked Function Expression
		// totalStudent := 30
		// sudentPresent := 25
	
		// var newPresent = func(total int, present int) float32 {
		// 	result := (float32(present) / float32(total)) * 100
	
		// 	return result
	
		// }(totalStudent, sudentPresent) // fungsinya langsung di-invoke saat deklarasi
	
		// fmt.Println("Attendance percentage :", newPresent)	

		//RECURSIVE
		// fmt.Println("Countdown Starts:")
		// // pemanggilan fungsi `countDown()`
		// countDown(3)
		
		// Factorial of a number using Go Recursion
		// var number = 3
		// var result = factorial (number)
		// fmt.Println("The factorial of 3 is", result)

		// Difference between Recursion and Iteration
		loop := factorialLoop(20)
		fmt.Println(loop)
		recursive := factorialRecursive(20)
		fmt.Println(recursive)

	
}

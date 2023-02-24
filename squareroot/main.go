package main

import "fmt"

// wap to find out squareroot of given number



func sqrt(num float64) float64 {
	if num < 0 {
		return -1
	}

	x := num
	y := 1.0
	accuracy := 0.000001

	for x-y > accuracy {
		x = (x + y) / 2
		y = num / x
	}

	return x
}

func main() {
	var num float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result := sqrt(num)

	if result == -1 {
		fmt.Println("Cannot find square root of a negative number")
		return
	}

	fmt.Printf("Square root of %.2f is %.2f\n", num, result)
}


func main() {
	var num float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Cannot find square root of a negative number")
		return
	}

	result := math.Sqrt(num)
	fmt.Printf("Square root of %.2f is %.2f\n", num, result)
}
In this program, we first ask the user to enter a number. Then, we check if the number is negative or not. If it is negative, we print an error message and return from the program. Otherwise, we use the math.Sqrt() function to calculate the square root of the number and print the result to the console.

Note that we've used the fmt package to print messages to the console and the math package to calculate the square root. Also, we've used the Scan() function from the fmt package to read the user's input from the console.




dikshathoke15@gmail.com
WAP for getting a square root of a given number in go without math packge
Sure, here's a simple Go program to calculate the square root of a given number without using the math package:

go
Copy code
package main

import (
	"fmt"
)

func sqrt(num float64) float64 {
	if num < 0 {
		return -1
	}

	x := num
	y := 1.0
	accuracy := 0.000001

	for x-y > accuracy {
		x = (x + y) / 2
		y = num / x
	}

	return x
}

// func main() {
// 	var num float64
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&num)

// 	result := sqrt(num)

// 	if result == -1 {
// 		fmt.Println("Cannot find square root of a negative number")
// 		return
// 	}

// 	fmt.Printf("Square root of %.2f is %.2f\n", num, result)
// }





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

	if num < 0 {
		fmt.Println("Cannot find square root of a negative number")
		return
	}

	result := math.Sqrt(num)
	fmt.Printf("Square root of %.2f is %.2f\n", num, result)
}


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





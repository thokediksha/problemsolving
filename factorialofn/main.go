package main

// Print Factorial of n
// E.g n=4, output=24 (1*2*3*4=24)

import "fmt"

func main() {
	 n := 4
	 fact := 1

	if n < 0 {
		fmt.Print("\nFactorial of a negative number is not possible")
	}
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	fmt.Printf("Factorial of %d is %d", n, fact)
}




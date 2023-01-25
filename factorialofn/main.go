package main

// Print Factorial of n
// E.g n=4, output=24 (1*2*3*4=24)

import "fmt"

func factorial(n int) int {
	if n == 1 || n == 0 {
      return n
	} else {
		return n * factorial(n-1)
	}
}

func main(){
  fmt.Println(factorial(4))
}
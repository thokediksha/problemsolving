package main


// Check whether number is a prime number or not
// E.g. n=6, output=True 

import (
	"fmt"
	"math"
)

func prime(n int) {
	if n < 2 {
		fmt.Println("number must be greater than 2")
		return
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			fmt.Println("not prime")
			return 
		}
	}
	fmt.Println("prime number")
	return 
}

func main() {
	prime(2)
}

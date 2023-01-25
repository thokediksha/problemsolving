package main 

// Print sum of first n numbers
// E.g. n=5, output=15 (1+2+3+4+5=15)

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
        return  n + sum(n - 1)
	} 
}

func main() {
    num := 5
	fmt.Printf("sum of first %d numbers is %d", num, sum(num))
} 

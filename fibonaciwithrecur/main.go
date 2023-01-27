package main

// WAP to print Fibonacci series with recursion

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
 return fibonacci(n-1) + fibonacci(n-2)
	}}

func main() {
    n := 5
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
}
 

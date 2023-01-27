package main

import "fmt"

    func main() {
		n := 10
		first := 0
		second := 1
		fmt.Print("Fibonacci series: ")
	
		for i := 0; i < n; i++ {
			if i <= 1 {
				fmt.Print(i, " ")
				continue
			}
			next := first + second
			fmt.Print(next, " ")
			first = second
			second = next
		}
	}
      

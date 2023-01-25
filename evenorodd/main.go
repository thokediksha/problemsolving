package main

// Print list of odd/even numbers from 1-n, where n is input.
// E.g. 
// n = 10
// even_nums = [2, 4, 6, 8, 10]
// odd_nums = [1, 3, 5, 7, 9]

import "fmt"

func even(n int) {
	// n := 10
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print(i, "\t")
		}
	}
}

func odd(n int) {
	// n := 10
	for i := 1; i <= n; i = i + 2 {
		fmt.Print(i, "\t")
	}
}

func main() {
	odd(5)
	fmt.Println()
	even(5)
}

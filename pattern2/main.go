package main


// WAP to print following pattern
// 1
// 1 2
// 1 2 3

import "fmt"

func main() {
var rows = 3
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
}
}
package main

// Find prime numbers from 1 ....n

import "fmt"

// func main() {
// 	n := 11
// 	for i := 2; i <= n; i++ {
// 		isPrime := true
// 		for j := 2; j < i; j++ {
// 			if i%j == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }


func main(){
n := 11
for i := 2; i <= n; i++ {
	isPrime := true
	for j := 2; j < i; j++{
		if i%j == 0 {
			isPrime = false
		}
	}
	if isPrime {
	fmt.Print(i, " ")
	}
}
}
package main

import "fmt"

    // func main() {
	// 	n := 10
	// 	first := 0
	// 	second := 1
	// 	fmt.Print("Fibonacci series: ")
	
	// 	for i := 0; i < n; i++ {
	// 		if i <= 1 {
	// 			fmt.Print(i, " ")
	// 			continue
	// 		}
	// 		next := first + second
	// 		fmt.Print(next, " ")
	// 		first = second
	// 		second = next
	// 	}
	// }

	func main(){
		fibo(10)
	 }
	 
	 func fibo(n int) {
		 n1 := 0
		 n2 := 1
		 n3 := 0
	 
		 for i:= 0; i <= n; i++ {
			 n3 = n1 + n2
			 fmt.Print(n1, " ")
			 n1 = n2
			 n2 = n3
		 }
	 }
      


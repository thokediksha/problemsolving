package main

// Swap Values without using any variable a=10, b=12

import "fmt"

// func main() {
//         a := 10
//         b := 12
//     a , b = b ,a 
//    fmt.Println(a,b)
// }


func main(){
	var row = 3
	for i := 0; i <= row; i++ {
		for j := 0; j < row-i; j++ {
		   fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
		   fmt.Print("*")
		}
		fmt.Println()
	 }
}
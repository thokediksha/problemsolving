package main


// // Check whether number is a prime number or not
// // E.g. n=6, output=True 

// import (
// 	"fmt"
// 	"math"
// )

// // func prime(n int) {
// // 	if n < 2 {
// // 		fmt.Println("number must be greater than 2")
// // 		return
// // 	}
// // 	sq_root := int(math.Sqrt(float64(n)))
// // 	for i := 2; i <= sq_root; i++ {
// // 		if n%i == 0 {
// // 			fmt.Println("not prime")
// // 			return q	
// // 		}
// // 	}
// // 	fmt.Println("prime number")
// // 	return 
// // }



// func prime(n1, n2 int){
//     if n1 < 2 || n2 < 2 {
//         fmt.Println("must be > 2")
//          return
//     }
//     for n1 <= n2 {
//         isPrime := true 
//         for i := 2; i <= int(math.Sqrt(float64(n1))); i++{
//             if n1 % i == 0 {
//                 isPrime = false
//                 break
//             }
//         }
//         if isPrime {
//             fmt.Printf("%d ", n1)
//         }
//        n1++
//     }
//     fmt.Println()
// }

// func main() {
// 	prime(5,11)
// }

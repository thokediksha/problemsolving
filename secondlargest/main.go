package main

import (
    "fmt"
    // "math"
)

//  WAP to find Second largest element in an array : arr[] = [12, 35, 1, 10, 34, 1, 35], without sorting, without using any built-in methods and without deleting duplicate elements. What will be the time complexity?
//  time complexity O(n)

// -1e9 is -1000000000.0; the minus sign applies to the number itself.
// func secondLargest(arr []int) int {
// 	second := int(-1e9)  
//     first := second
//     for i := 0; i < len(arr); i++ {
//         if arr[i] > first {
// 			second = first
//             first = arr[i]
//         } else if arr[i] > second && arr[i] != first {
//             second = arr[i]
//         }
//     }
//     return second
// }

// func main(){
// 	arr := []int {1,2,3,4,5,100,99}
// 	fmt.Println(secondLargest(arr))
// }



// func main(){
//     var l1, l2 = 0, 0
//      a  := []int {1,2,3,4,5,100,99}

//     l1 = a[0]
//     for i := 1; i <= 6; i++{
//         if l1 < a[i] {
//            l2 = l1
//            l1 = a[i]
//         } else if l2 < a[i] {
//             l2 = a[i]
//         }
//     }
//     fmt.Println("large", l2)
// }

// func main() {
//     n := 4
//     var j int
//     for i:=1; i<=n; i++{
//         for j=1; j<i; j++{
//            fmt.Printf(" ")
//         }
//         for k:=i; k<n; k++{
//             if i == 2 {
//                 fmt.Print("*")
//             } else {
//                 fmt.Print("", k)
//             }
//      }
//      fmt.Println()
// }
// }

func main(){
    var l1 = 0
    var l2 = 0
    a := []int {1,2,3,4,6,54,43}

    l1 = a[0]
    for i:= 1; i<=6; i++{
        if l1 < a[i] {
            l2 = l1
            l1 = a[i]
        }else if l2 < a[i]{
            l2 = a[i]
        }
    }
    fmt.Println(l1)
}




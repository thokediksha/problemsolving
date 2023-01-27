package main

import "fmt"

func main(){

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 101}
    fmt.Print("Even numbers in the list: ")
    for _, num := range nums {
        if num % 2 == 0 {
            fmt.Print(num, " ")
        }
    }
}



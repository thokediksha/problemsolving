package main

// WAP to print following pattern (use while loop):
// 3 2 1
//    2 1
//       1


import "fmt"

func main(){
	i := 3
    for i > 0 {
        for j := i; j < 4; j++ {
            fmt.Print("  ")
        }
        for k := i; k > 0; k-- {
            fmt.Print(k," ")
        }
        fmt.Println()
        i--
    }
}



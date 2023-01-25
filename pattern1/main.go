package main


// n = 19
// 1 2 3 4 5 * * * * * 11 12 13 14 15 * * * *

// Here sets are divided by 5, first set prints numbers where next set prints * so on and so forth

// If n=12, output: 1 2 3 4 5 * * * * * 11 12
// If n=4, output: 1 2 3 4


import "fmt"

func main() {
    n := 4
    i := 1
    for i <= n {
        if i <= 5 {
            fmt.Print(i, " ")
        } else if i <= 10 {
            fmt.Print("* ")
        } else if i <= 15 {
			fmt.Print(i, " ")
		} else if i <= 20 {
			fmt.Print("* ")
		} 
        i++
    }
}




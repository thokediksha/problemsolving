package main

// Given a string of 0s and 1s print the value of string, where val(0)=1, val(1)=2
// Input: 00010110
// Output: 11


import 
	"fmt"

func getVal(str string) int {
    val := 0
    for _, c := range str {
        if c == '0' {
            val = val + 1
        } else {
            val = val + 2 
        }
    }
    return val
}

func main() {
    s := "00010110"
    fmt.Println(getVal(s))
}



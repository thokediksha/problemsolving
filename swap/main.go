package main

// Swap Values without using any variable a=10, b=12

import "fmt"

func main() {
        a := 10
        b := 12
    a , b = b ,a 
   fmt.Println(a,b)
}


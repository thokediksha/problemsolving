package main

// Generate a dictionary which stores the count of appearance of character in string
// Input: aabbcdddededfg
// Output= {“a”: 2, “b”: 2, “c”: 1, “d”:5, “e”:1, “f”:1, “g”:1}

import (
	"fmt"
	"strings"
)

func main(){
	str := "aabbcdddededfg"
	res1 := strings.Count(str, "a")
	res2 := strings.Count(str, "b")
	res3 := strings.Count(str, "c")
	res4 := strings.Count(str, "d")
	res5 := strings.Count(str, "e")
	res6 := strings.Count(str, "f")
	res7 := strings.Count(str, "g")
	fmt.Println("{"+"a:",res1,","+"b:",res2,","+"c:",res3,","+"d:",res4,","+"e:",res5,","+"f:",res6,","+"g:",res7,"}")
}

// func main() {
//     input := "aabbcdddededfg"
//     charCount := map[rune]int{}
//     for _, char := range input {
//         charCount[char]++
//     }
//     output := make(map[string]int)
//     for key, value := range charCount {
//         output[string(key)] = value
//     }
//     fmt.Println(output)
// }
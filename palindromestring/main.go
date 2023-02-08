package main

import (
	"fmt"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main(){
	str := "hello"
	strREV := reverse(str)
	fmt.Println(str)
	fmt.Println(strREV)
	if str == strREV {
		fmt.Println("isPalindrome")
	} else {
        fmt.Println("not isPalindrome")
	}
}


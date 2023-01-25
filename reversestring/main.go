package main

// Find the reverse of the string

import "fmt"

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main(){
	str := "DIKSHA"
	strREV := reverse(str)
	fmt.Println(str)
	fmt.Println(strREV)
}
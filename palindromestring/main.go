package main

import (
	"fmt"
	"strings"
)

func main() {
	originalString := "hello"

    var reverseString string = ""
    var length = len(originalString)

    for i := length - 1; i >= 0; i -- {
        reverseString = reverseString + string(originalString[i])
    }

    if strings.ToLower(originalString) == strings.ToLower(reverseString) {
        fmt.Println("The given string is Palindrome");
    } else {
        fmt.Println("The given string is NOT a Palindrome");
    }
}
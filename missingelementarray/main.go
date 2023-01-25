package main

// . WAP to find missing elements from the array

import "fmt"

func missingnumber(arr []int, n int) int {
	sum := 0
	nsum := 0
	for i := 0; i < n; i++ {
		sum += arr[i] // sum variable will contain the sum fo all integers of an array.
	}
	nsum = ((n + 1) * (n + 2)) / 2 // we use formula for obtaining sum of first n numbers.
	return (nsum - sum)            // return the missing number.
}

func main() {
	arr := []int{1,2,3,4,5,7,8} // declaring and initialising the array having n-1(8) integers.
	size1 := len(arr)
	missing_no := missingnumber(arr, size1)
	fmt.Printf("The missing number in the array of size %d is %d", size1+1, missing_no) // Printing the missing number
}

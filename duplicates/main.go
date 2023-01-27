package main

// wap to find out duplicate elements in an array

import "fmt"

func duplicate(arr []int) []int {

	var duplicates []int

	numfreq := make(map[int]int)

	for _, num := range arr {
		numfreq[num]++
	}

	for num, freq := range numfreq {
		if freq > 1 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}

func main() {
	arr := []int{1, 1, 2, 3, 4, 5, 6, 7, 7}
	duplicates := duplicate(arr)
	fmt.Println(duplicates)
}

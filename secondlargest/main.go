package main

import "fmt"

//  WAP to find Second largest element in an array : arr[] = [12, 35, 1, 10, 34, 1, 35], without sorting, without using any built-in methods and without deleting duplicate elements. What will be the time complexity?
//  time complexity O(n)

// -1e9 is -1000000000.0; the minus sign applies to the number itself.
func secondLargest(arr []int) int {
	second := int(-1e9)  
    first := second
    for i := 0; i < len(arr); i++ {
        if arr[i] > first {
			second = first
            first = arr[i]
        } else if arr[i] > second && arr[i] != first {
            second = arr[i]
        }
    }
    return second
}

func main(){
	arr := []int {1,2,3,4,5,100,99}
	fmt.Println(secondLargest(arr))
}


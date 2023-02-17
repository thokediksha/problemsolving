package main

// Find the intersection of two lists
// E.g List1=[1, 2, 3, 4, 5] List2=[3, 4, 6] Intersection=[3, 4]

import (
	"fmt"
	// "math"
)

// func intersection(a, b []int) (c []int){
//     m := make(map[int]bool) 
// 	for _, item:= range a {
// 		m[item] = true
// 	}
// 	for _, item:= range b {
// 		if _, ok := m[item]; ok {
// 			c = append(c,item)
// 		}
// 	}
// 	return
// }

func main(){
	l1 := 0
	l2 := 0
	a := []int {1,2,3,4,100}

	l1 = a[0]
	for i := 1; i <= 4; i++{
		if l1 < a[i] {
			l2 = l1
			l1 = a[i]
		} else if l2 < a[i]{
			l2 = a[i]
		}
	}
	fmt.Println(l1)
}





package main

// Find the intersection of two lists
// E.g List1=[1, 2, 3, 4, 5] List2=[3, 4, 6] Intersection=[3, 4]

import (
	"fmt"
	"math"
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

// func main(){
// 	var a = []int{1, 2, 3, 4, 5}
// 	var b = []int{3, 4, 6}
// 	fmt.Println(intersection(a,b))
// }



func inter(a,b []int) (c []int) {
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
         c := append(c, item)
		}
	}
}


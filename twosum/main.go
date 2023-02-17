package main

import "fmt"

func twosum(nums []int, target int) []int {
	n := len(nums)
	for i:=0; i<n-1; i++{
		for j:=i+1; j<n; j++{
			if nums[i] + nums[j] == target {
				return []int {i,j}
			}
		} 
	}
	return []int {}
}

func main() {
	nums := []int {1,2,3,4,5}
	target := 5
	c := twosum(nums, target)
	fmt.Println(c)
}
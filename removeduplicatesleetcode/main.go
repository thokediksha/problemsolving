package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	  
	  if n == 0 {
		  
		  return 0
	  }
	  
   index := 1

	  for i := 1; i < n; i++ {
		  
		  if nums[i] != nums[i-1] {
					  
			  nums[index] = nums[i]
			  index++
		  }
	  }
	  
	  
	  return index
}

func main(){
	nums := []int {1,1,2}
	c := removeDuplicates(nums)
	fmt.Println(c)
}

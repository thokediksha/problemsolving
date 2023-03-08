package main

// wap to find out pairs whose sum is 100

import "fmt"

func findPairs(arr []int, sum int) [][]int {
    var pairs [][]int
    n := make(map[int]bool)
    for _, val := range arr {
        diff := sum - val
        if n[diff] {
            pairs = append(pairs, []int{diff, val})
        }
        n[val] = true
    }
    return pairs
}

func main() {
    arr := []int{80, 60, 10, 50, 30, 100, 0, 40, 50}
    sum := 100
    pairs := findPairs(arr, sum)
    fmt.Println("Pairs whose sum is", sum, ":", pairs)
}
package main

// logic for anagram program with its time complexity. (for large strings)

import "fmt"


func isAnagram(s1,s2 string) bool {
	fre := make(map[rune]int)

	for _, r := range s1{
		fre[r]++
	}
	for _, r := range s2{
		fre[r]--
	}
	for _, count := range fre{
		if count  != 0 {
			return false
		}
	}
	return true
}

func main() {
    s1 := "madam"
    s2 := "damma"
    if isAnagram(s1, s2) {
        fmt.Println("The strings are anagrams.")
    } else {
        fmt.Println("The strings are not anagrams.")
    }
}

// Time complexity of this program is O(n).










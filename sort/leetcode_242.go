package main

import (
	"sort"
)

//func main() {
//	s, t := "anagram", "nagaram"
//	result := isAnagram(s, t)
//	fmt.Println(result)
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1, t1 := []rune(s), []rune(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	for i := 0; i < len(s); i++ {
		if s1[i] != t1[i] {
			return false
		}
	}
	return true
}

package main

import (
	"sort"
)

//func main() {
//	s := "abb"
//	answer := longestPalindrome(s)
//	fmt.Println("=====")
//	fmt.Println(answer)
//}

func longestPalindrome(s string) string {
	candidate := make([]string, 0)
	stringSlice := []rune(s)
	reversedString := reverse(s)
	stringLength := len(stringSlice)
	for i := 1; i <= stringLength; i++ {
		answer := make([]rune, 0)
		for k := 0; k < i; k++ {
			if stringSlice[k] == reversedString[stringLength-i+k] {
				answer = append(answer, stringSlice[k])
			} else {
				if len(answer) == 0 {
					continue
				}
				candidate = append(candidate, string(answer))
				answer = []rune{stringSlice[k]}
			}
		}
		if len(answer) == 0 {
			continue
		}
		candidate = append(candidate, string(answer))
	}
	sort.Slice(candidate, func(i, j int) bool {
		return len(candidate[i]) > len(candidate[j])
	})
	return candidate[0]
}

func reverse(s string) []rune {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

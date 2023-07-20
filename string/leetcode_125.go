package main

import (
	"strings"
)

//func main() {
//	s := "0P"
//	ans := isPalindrome(s)
//	fmt.Println(ans)
//}

func isPalindrome(s string) bool {
	sRune := []rune(strings.ToLower(s))
	i, j := 0, len(sRune)-1
	for i < j {
		if isAlphaNumeric(sRune[i]) {
			if isAlphaNumeric(sRune[j]) {
				if sRune[i] != sRune[j] {
					return false
				} else {
					i++
					j--
				}
			} else {
				j--
			}
		} else {
			i++
			continue
		}
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	if ('a' <= r && r <= 'z') || ('0' <= r && r <= '9') {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "0P"
	ans := isPalindrome(s)
	fmt.Println(ans)
}

func isPalindrome(s string) bool {
	str := []rune(strings.ToLower(s))
	res := true
	start := 0
	end := len(s) - 1
	lastIndex := end
	for {
		for {
			if start < lastIndex && !isAlphabet(str[start]) {
				start++
				continue
			}
			break
		}
		for {
			if end > 0 && !isAlphabet(str[end]) {
				end--
				continue
			}
			break
		}
		if start >= end {
			break
		}
		if str[start] == str[end] {
			start++
			end--
			if start >= end {
				break
			}
		} else {
			res = false
			break
		}
	}
	return res
}

func isAlphabet(char rune) bool {
	// LowerCase 97 - 122
	// number 48 - 57
	if (char >= 97 && char <= 122) || (char >= 48 && char <= 57) {
		return true
	} else {
		return false
	}
}

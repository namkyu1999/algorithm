package main

import (
	"sort"
)

//func main() {
//	s := "aacabdkacaa"
//	answer := longestPalindrome(s)
//	fmt.Println(answer)
//}

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		palindromeOdd := expandAroundCenter(s, i, i)
		palindromeEven := expandAroundCenter(s, i, i+1)

		if len(palindromeOdd) > len(longest) {
			longest = palindromeOdd
		}
		if len(palindromeEven) > len(longest) {
			longest = palindromeEven
		}
	}
	return longest
}

func expandAroundCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome2(s string) string {
	candidate, stringSlice, size := make([]string, 0), []rune(s), len(s)
	evenIndex, oddIndex := 0, 0
	if s == string(reverse(s)) {
		return s
	}
	if size == 2 {
		if stringSlice[0] == stringSlice[1] {
			return s
		} else {
			return string(stringSlice[0])
		}
	}
	for {
		if evenIndex <= size-2 {
			scope := evenIndex
			for {
				if checkInSlice(scope, size) && checkInSlice(scope+(evenIndex-scope)*2+1, size) {
					if stringSlice[scope] == stringSlice[scope+(evenIndex-scope)*2+1] {
						candidate = append(candidate, string(stringSlice[scope:scope+(evenIndex-scope)*2+2]))
					} else {
						break
					}
					scope--
				} else {
					break
				}
			}
			evenIndex += 1
		}
		if oddIndex <= size-3 {
			scope := oddIndex
			for {
				if checkInSlice(scope, size) && checkInSlice(scope+(oddIndex-scope)*2+2, size) {
					if stringSlice[scope] == stringSlice[scope+(oddIndex-scope)*2+2] {
						candidate = append(candidate, string(stringSlice[scope:scope+(oddIndex-scope)*2+3]))
					} else {
						break
					}
					scope--
				} else {
					break
				}
			}
			oddIndex += 1
		}
		if evenIndex > size-2 && oddIndex > size-3 {
			break
		}
	}
	if len(candidate) == 0 {
		return string(stringSlice[0])
	}
	sort.Slice(candidate, func(i, j int) bool {
		return len(candidate[i]) > len(candidate[j])
	})
	return candidate[0]
}

func checkInSlice(current, size int) bool {
	if current >= 0 && current < size {
		return true
	} else {
		return false
	}
}

func reverse(s string) []rune {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

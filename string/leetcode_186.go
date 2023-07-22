package main

import "fmt"

func main() {
	str := "the sky is blue"
	reverseWords([]byte(str))
}

func reverseWords(s []byte) {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	start, end := 0, 0
	for end < len(s) {
		for ; end < len(s); end++ {
			if s[end] == ' ' {
				break
			}
		}
		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		start, end = end+1, end+1
	}
	fmt.Println(string(s))
}

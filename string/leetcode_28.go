package main

import (
	"strings"
)

//func main() {
//	hayStack := "aaaaa"
//	needle := "bba"
//	result := strStr(hayStack, needle)
//	fmt.Println(result)
//}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

package main

import (
	"regexp"
	"strings"
)

//func main() {
//	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
//	banned := []string{"hit"}
//	mostCommon := mostCommonWord(paragraph, banned)
//	fmt.Println(mostCommon)
//}

func mostCommonWord(paragraph string, banned []string) string {
	max, maxIndex := "", 0
	paragraph = strings.ToLower(paragraph)
	paragraph = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(paragraph, " ")
	hash := map[string]int{}
	for _, v := range strings.Split(paragraph, " ") {
		if v == "" {
			continue
		}
		word := v
		hash[word] += 1
		if hash[word] > maxIndex && !stringContains(banned, word) {
			max = word
			maxIndex = hash[word]
		}
	}
	return max
}

func stringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

package main

import (
	"sort"
	"strconv"
	"strings"
)

//func main() {
//	stringArray := []string{"dig1 8 1 5 1", "let1 art zero can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
//	res := reorderLogFiles(stringArray)
//	fmt.Println(res)
//}

func reorderLogFiles(logs []string) []string {
	var numberLogs []string
	var stringLogs []string

	for _, v := range logs {
		flag := strings.Split(v, " ")[1]
		if _, err := strconv.ParseComplex(flag, 128); err == nil {
			numberLogs = append(numberLogs, v)
		} else {
			stringLogs = append(stringLogs, v)
		}
	}
	orderAsDictionary(stringLogs)
	return append(stringLogs, numberLogs...)
}

func orderAsDictionary(chunks []string) {
	hash := make(map[string]string)
	keyHash := make(map[string][]string)
	compareChunks := make([]string, 0)
	for i, v := range chunks {
		tmp := strings.Split(v, " ")
		newKey := tmp[0]
		newString := strings.Join(tmp[1:], " ")
		keyHash[newString] = append(keyHash[newString], newKey)
		hash[newString] = chunks[i]
		compareChunks = append(compareChunks, newString)
	}
	sort.Strings(compareChunks)
	for i := 0; i < len(chunks); i++ {
		value := compareChunks[i]
		duplicatedCount := len(keyHash[value])
		if duplicatedCount > 1 {
			sort.Strings(keyHash[value])
			for k, v := range keyHash[value] {
				chunks[i+k] = v + " " + value
			}
			i += duplicatedCount
			continue
		} else {
			chunks[i] = hash[value]
		}
	}
}

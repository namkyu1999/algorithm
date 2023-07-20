package main

import (
	"sort"
	"strings"
)

//func main() {
//	stringArray := []string{"dig1 8 1 5 1", "let1 art zero can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
//	res := reorderLogFiles(stringArray)
//	fmt.Println(res)
//}

func reorderLogFiles(logs []string) []string {
	var letterLogs []string
	i := 0
	for i < len(logs) {
		char := rune(strings.Split(logs[i], " ")[1][0])
		if '0' <= char && char <= '9' {
			// digit
			i++
			continue
		} else {
			// letter
			letterLogs = append(letterLogs, logs[i])
			logs = append(logs[:i], logs[i+1:]...)
		}
	}
	sort.Slice(letterLogs, func(i, j int) bool {
		iStr, jStr := strings.SplitN(letterLogs[i], " ", 2), strings.SplitN(letterLogs[j], " ", 2)
		flag := strings.Compare(iStr[1], jStr[1])
		if flag == -1 {
			return true
		} else if flag == 0 {
			flag2 := strings.Compare(iStr[0], jStr[0])
			if flag2 == -1 {
				return true
			}
			return false
		}
		return false
	})
	return append(letterLogs, logs...)
}

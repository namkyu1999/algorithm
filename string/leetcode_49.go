package main

import (
	"sort"
)

//func main() {
//	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//	answer := groupAnagrams(strs)
//	fmt.Println(answer)
//}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	hash := map[string][]string{}
	for _, v := range strs {
		runeString := []rune(v)
		sort.Slice(runeString, func(i, j int) bool {
			return runeString[i] < runeString[j]
		})
		hash[string(runeString)] = append(hash[string(runeString)], v)
	}
	for _, v := range hash {
		result = append(result, v)
	}
	return result
}

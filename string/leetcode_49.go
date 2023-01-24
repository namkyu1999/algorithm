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
	hash := make(map[string][]string, 0)
	for _, v := range strs {
		sorted := SortStringByCharacter(v)
		if _, ok := hash[sorted]; ok {
			hash[sorted] = append(hash[sorted], v)
		} else {
			hash[sorted] = []string{v}
		}
	}
	result := make([][]string, 0)
	for _, v := range hash {
		result = append(result, v)
	}
	return result
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

//func groupAnagrams(strs []string) [][]string {
//	answer := make([][]string, 0)
//	answer = append(answer, strs)
//	i := 0
//	for {
//		if len(answer[i]) > 1 {
//			isExpanded := false
//			key := answer[i][0]
//			j := 1
//			for {
//				isAnagram := checkAnagram(key, answer[i][j])
//				if isAnagram {
//					if len(answer[i]) > j+1 {
//						j++
//					} else {
//						break
//					}
//				} else {
//					if isExpanded {
//						answer[i+1] = append(answer[i+1], answer[i][j])
//					} else {
//						isExpanded = true
//						answer = append(answer, []string{answer[i][j]})
//					}
//					if j == len(answer[i])-1 {
//						answer[i] = answer[i][:j]
//						break
//					} else {
//						answer[i] = deleteWithIndex(answer[i], j)
//					}
//				}
//			}
//		}
//		if len(answer) > i+1 {
//			i++
//		} else {
//			break
//		}
//	}
//	return answer
//}

//func deleteWithIndex(slice []string, index int) []string {
//	return append(slice[:index], slice[index+1:]...)
//}

//func checkAnagram(a string, b string) bool {
//	if len(a) != len(b) {
//		return false
//	}
//	a = SortStringByCharacter(a)
//	b = SortStringByCharacter(b)
//
//	return a == b
//}

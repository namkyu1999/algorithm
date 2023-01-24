package main

import (
	"regexp"
	"sort"
	"strings"
)

//func main() {
//	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
//	banned := []string{"hit"}
//	mostCommon := mostCommonWord(paragraph, banned)
//	fmt.Println(mostCommon)
//}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	re, _ := regexp.Compile(`[!?',;.]`)
	paragraph = re.ReplaceAllString(paragraph, " ")
	stringArray := strings.Split(paragraph, " ")
	stringMap := make(map[string]int)
	for _, v := range stringArray {
		if v == "" {
			continue
		}
		stringMap[v] += 1
	}

	p := make(PairList, len(stringMap))

	i := 0

	for k, v := range stringMap {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	for _, v := range p {
		isBanned := contains(v.Key, banned)
		if !isBanned {
			return v.Key
		}
	}
	return ""
}

func contains(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

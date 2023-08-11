package main

import (
	"sort"
	"strings"
)

//func main() {
//	tickets := [][]string{
//		{"JFK", "SFO"},
//		{"JFK", "ATL"},
//		{"SFO", "ATL"},
//		{"ATL", "JFK"},
//		{"ATL", "SFO"},
//	}
//	result := findItinerary(tickets)
//	fmt.Println(result)
//}

func findItinerary(tickets [][]string) []string {
	type flight struct {
		to   string
		used bool
	}
	hash := map[string][]flight{}
	for _, v := range tickets {
		hash[v[0]] = append(hash[v[0]], flight{to: v[1]})
		sort.Slice(hash[v[0]], func(i, j int) bool {
			return strings.Compare(hash[v[0]][i].to, hash[v[0]][j].to) == -1
		})
	}
	var dfs func(routes *[]string, left int) int
	dfs = func(routes *[]string, left int) int {
		if left == 0 {
			return 0
		}
		current := (*routes)[len(*routes)-1]
		for i := 0; i < len(hash[current]); i++ {
			if hash[current][i].used {
				continue
			}
			hash[current][i].used = true
			*routes = append(*routes, hash[current][i].to)
			result := dfs(routes, left-1)
			if result == 0 {
				return 0
			}
			*routes = (*routes)[:len(*routes)-1]
			hash[current][i].used = false
		}
		return left
	}
	routes := &[]string{"JFK"}
	dfs(routes, len(tickets))
	return *routes
}

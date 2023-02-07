package main

import (
	"sort"
)

//func main() {
//	nums := []int{1, 1, 1, 2, 2, 3}
//	k := 2
//	result := topKFrequent(nums, k)
//	fmt.Println(result)
//}

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v] += 1
	}
	keys := make([]int, 0, len(hash))
	for key := range hash {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return hash[keys[i]] > hash[keys[j]]
	})
	return keys[:k]
}

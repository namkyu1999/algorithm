package main

import (
	"math"
)

//func main() {
//	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
//	n, k := 4, 2
//	result := networkDelayTime(times, n, k)
//	fmt.Println(result)
//}

func networkDelayTime(times [][]int, n int, k int) int {
	table, visited := make([]int, n+1), make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = math.MaxInt
	}
	table[k] = 0
	shortest, count := k, 1

	hash := map[int][][]int{}
	for _, time := range times {
		hash[time[0]] = append(hash[time[0]], time[1:])
	}

	for count < n {
		visited[shortest] = true
		for _, v := range hash[shortest] {
			if table[v[0]] > v[1]+table[shortest] {
				table[v[0]] = v[1] + table[shortest]
			}
		}

		candidate := 0
		for i, v := range table {
			if v < table[candidate] && !visited[i] {
				candidate = i
			}
		}
		shortest = candidate
		count++
	}
	result := 0
	for _, v := range table[1:] {
		if result < v {
			result = v
		}
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}

package main

import (
	"fmt"
)

//func main() {
//	graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
//	result := allPathsSourceTarget(graph)
//	fmt.Println(result)
//}

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int

	var dfs func([]int, int)
	dfs = func(path []int, node int) {
		fmt.Printf("%p\n", path)
		path = append(path, node)
		if node == len(graph)-1 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			res = append(res, newPath)
			return
		}
		for _, neighbor := range graph[node] {
			dfs(path, neighbor)
		}
	}
	dfs([]int{}, 0)
	return res
}

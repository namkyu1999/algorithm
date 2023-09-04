package main

//func main() {
//	graph := [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {3, 4}, {4}, {}}
//	result := eventualSafeNodes(graph)
//	fmt.Println(result)
//}

func eventualSafeNodes(graph [][]int) []int {
	var isCycle func(current int) bool
	visited := make([]bool, len(graph))
	memoization := map[int]bool{}
	isCycle = func(current int) bool {
		if _, ok := memoization[current]; ok {
			return memoization[current]
		}
		visited[current] = true
		for _, v := range graph[current] {
			if visited[v] || isCycle(v) {
				visited[current] = false
				memoization[current] = true
				return true
			}
		}
		visited[current] = false
		memoization[current] = false
		return false
	}
	var result []int
	for i := 0; i < len(graph); i++ {
		if !isCycle(i) {
			result = append(result, i)
		}
	}
	return result
}

package main

//func main() {
//	graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
//	result := allPathsSourceTarget(graph)
//	fmt.Println(result)
//}

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	queue := [][]int{{0}}
	destination := len(graph) - 1
	for len(queue) > 0 {
		current := queue[0]
		element := current[len(current)-1]
		queue = queue[1:]
		if element == destination {
			result = append(result, current)
			continue
		}
		for _, v := range graph[element] {
			next := make([]int, len(current)+1)
			copy(next, current)
			next[len(current)] = v
			queue = append(queue, next)
		}
	}
	return result
}

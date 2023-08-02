package main

//func main() {
//	n := 10
//	edges := [][]int{
//		{0, 7},
//		{0, 8},
//		{6, 1},
//		{2, 0},
//		{0, 4},
//		{5, 8},
//		{4, 7},
//		{1, 3},
//		{3, 5},
//		{6, 5},
//	}
//	source := 7
//	destination := 5
//	result := validPath(n, edges, source, destination)
//	fmt.Println(result)
//}

func validPath(n int, edges [][]int, source int, destination int) bool {
	hash := map[int][]int{}
	for _, v := range edges {
		hash[v[0]] = append(hash[v[0]], v[1])
		hash[v[1]] = append(hash[v[1]], v[0])
	}
	visited := make([]bool, n)
	queue := []int{source}
	for len(queue) > 0 {
		current := queue[0]
		if current == destination {
			return true
		}
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		queue = append(queue, hash[current]...)
	}
	return false
}

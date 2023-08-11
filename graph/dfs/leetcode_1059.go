package main

//func main() {
//	n := 4
//	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
//	source := 0
//	destination := 3
//	result := leadsToDestination(n, edges, source, destination)
//	fmt.Println(result)
//}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	hash := make([][]int, n)
	for _, v := range edges {
		hash[v[0]] = append(hash[v[0]], v[1])
	}
	if len(hash[destination]) > 0 {
		return false
	}

	visited := make([]byte, n)

	var dfs func(int) bool
	dfs = func(src int) bool {
		if src == destination {
			return true
		}
		if len(hash[src]) == 0 {
			return false
		}
		if visited[src] > 0 {
			return false
		}
		visited[src] = 1
		for _, v := range hash[src] {
			if !dfs(v) {
				return false
			}
		}
		visited[src] = 0
		return true
	}
	return dfs(source)
}

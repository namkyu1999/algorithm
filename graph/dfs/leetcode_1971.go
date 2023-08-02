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
	result := false
	var dfs func(path []int, vertex int)
	dfs = func(path []int, vertex int) {
		if vertex == destination {
			result = true
			return
		}
		if visited[vertex] {
			return
		}
		visited[vertex] = true
		for _, v := range hash[vertex] {
			if result {
				return
			}
			dfs(append(path, vertex), v)
		}
		visited[vertex] = true
	}

	dfs([]int{}, source)
	return result
}

//func validPath_uf(n int, edges [][]int, source int, destination int) bool {
//	uf := InitUnionFound(n)
//	for _, v := range edges {
//		uf.Union(v[0], v[1])
//	}
//	return uf.Find(source) == uf.Find(destination)
//}
//
//type UnionFound struct {
//	rank []int
//	root []int
//}
//
//func InitUnionFound(size int) UnionFound {
//	rank, root := make([]int, size), make([]int, size)
//	for i := 0; i < size; i++ {
//		rank[i] = 1
//		root[i] = i
//	}
//	return UnionFound{rank: rank, root: root}
//}
//
//func (uf *UnionFound) Find(x int) int {
//	if uf.root[x] == x {
//		return x
//	}
//	uf.root[x] = uf.Find(uf.root[x])
//	return uf.root[x]
//}
//
//func (uf *UnionFound) Union(x, y int) {
//	rootX, rootY := uf.Find(x), uf.Find(y)
//	if rootY != rootX {
//		if uf.rank[rootX] > uf.rank[rootY] {
//			uf.root[rootY] = rootX
//		} else if uf.rank[rootX] < uf.rank[rootY] {
//			uf.root[rootX] = rootY
//		} else {
//			uf.rank[rootX]++
//			uf.root[rootY] = rootX
//		}
//	}
//}

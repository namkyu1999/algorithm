package main

//func main() {
//	n := 5
//	edges := [][]int{
//		{0, 1}, {1, 2}, {3, 4},
//	}
//	result := countComponents(n, edges)
//	fmt.Println(result)
//}
//
//type UnionFind struct {
//	root []int
//	rank []int
//}
//
//func Construct(size int) UnionFind {
//	root, rank := make([]int, size), make([]int, size)
//	for i := 0; i < size; i++ {
//		root[i] = i
//		rank[i] = 1
//	}
//	return UnionFind{
//		root: root,
//		rank: rank,
//	}
//}
//
//func (uf *UnionFind) Union(x, y int) {
//	rootX, rootY := uf.Find(x), uf.Find(y)
//	if rootX != rootY {
//		if uf.rank[rootX] > uf.rank[rootY] {
//			uf.root[rootY] = rootX
//		} else if uf.rank[rootY] > uf.rank[rootX] {
//			uf.root[rootX] = rootY
//		} else {
//			uf.root[rootY] = rootX
//			uf.rank[rootX]++
//		}
//	}
//}
//
//func (uf *UnionFind) Find(x int) int {
//	if x == uf.root[x] {
//		return x
//	}
//	uf.root[x] = uf.Find(uf.root[x])
//	return uf.root[x]
//}
//
//func countComponents(n int, edges [][]int) int {
//	uf := Construct(n)
//	for _, v := range edges {
//		uf.Union(v[0], v[1])
//	}
//	count := 0
//	for i, v := range uf.root {
//		if i == v {
//			count++
//		}
//	}
//	return count
//}

package main

//func main() {
//	isConnected := [][]int{
//		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
//	}
//	result := findCircleNum(isConnected)
//	fmt.Println(result)
//}
//
//type UnionFind struct {
//	root []int
//	rank []int
//}
//
//func InitUnionFind(size int) UnionFind {
//	uf := UnionFind{}
//	uf.root, uf.rank = make([]int, size), make([]int, size)
//	for i := 0; i < size; i++ {
//		uf.root[i] = i
//		uf.rank[i] = 1
//	}
//	return uf
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
//func findCircleNum(isConnected [][]int) int {
//	uf := InitUnionFind(len(isConnected))
//	for i, v := range isConnected {
//		for k := i + 1; k < len(v); k++ {
//			if v[k] == 1 {
//				uf.Union(i, k)
//			}
//		}
//	}
//	answer := 0
//	for i, v := range uf.root {
//		if i == v {
//			answer++
//		}
//	}
//	return answer
//}

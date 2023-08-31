package main

//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	n := 3
//	wells := []int{1, 2, 2}
//	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}
//	result := minCostToSupplyWater(n, wells, pipes)
//	fmt.Println(result)
//}
//
//func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
//	result := 0
//	tmp := make([]int, n+1)
//	copy(tmp[1:], wells)
//	wells = tmp
//	uf := InitUnionFound(n + 1)
//	sort.Slice(pipes, func(i, j int) bool {
//		return pipes[i][2] < pipes[j][2]
//	})
//	for len(pipes) > 0 {
//		nodeA, nodeB, cost := pipes[0][0], pipes[0][1], pipes[0][2]
//		pipes = pipes[1:]
//
//		rootA, rootB := uf.Find(nodeA), uf.Find(nodeB)
//
//		if rootA == rootB {
//			continue
//		}
//
//		if wells[rootA] > wells[rootB] {
//			if wells[rootA] > cost {
//				uf.Union(rootB, rootA)
//				result += cost
//			}
//		} else {
//			if wells[rootB] > cost {
//				uf.Union(rootA, rootB)
//				result += cost
//			}
//		}
//	}
//	set := map[int]bool{}
//	for _, v := range uf.root {
//		set[uf.Find(v)] = true
//	}
//	for k := range set {
//		result += wells[k]
//	}
//	return result
//}
//
//type UnionFound struct {
//	root []int
//}
//
//func InitUnionFound(n int) UnionFound {
//	root := make([]int, n)
//	for i := 0; i < n; i++ {
//		root[i] = i
//	}
//	return UnionFound{
//		root: root,
//	}
//}
//
//func (uf *UnionFound) Find(x int) int {
//	if x == uf.root[x] {
//		return x
//	}
//	uf.root[x] = uf.Find(uf.root[x])
//	return uf.root[x]
//}
//
//func (uf *UnionFound) Union(x, y int) {
//	rootX, rootY := uf.Find(x), uf.Find(y)
//	if rootX != rootY {
//		uf.root[rootY] = rootX
//	}
//}

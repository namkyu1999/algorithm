package main

import (
	"container/heap"
)

//func main() {
//	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
//	result := minCostConnectPoints(points)
//	fmt.Println(result)
//}

// kruskal
//func minCostConnectPoints(points [][]int) int {
//	result := 0
//	distance := map[int][][]int{}
//	for i, v := range points {
//		for k := i + 1; k < len(points); k++ {
//			dist := distAbs(v[0], points[k][0]) + distAbs(v[1], points[k][1])
//			distance[dist] = append(distance[dist], []int{i, k})
//		}
//	}
//
//	keys, i := make([]int, len(distance)), 0
//	for k := range distance {
//		keys[i] = k
//		i++
//	}
//
//	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
//
//	uf := InitUnionFound(len(points))
//	for _, path := range keys {
//		point := distance[path]
//		for _, p := range point {
//			union := uf.Union(p[0], p[1])
//			if union {
//				result += path
//			}
//		}
//	}
//	return result
//}
//
//func distAbs(a, b int) int {
//	result := a - b
//	if result > 0 {
//		return result
//	}
//	return -result
//}
//
//type UnionFound struct {
//	root []int
//	rank []int
//}
//
//func InitUnionFound(size int) UnionFound {
//	root, rank := make([]int, size), make([]int, size)
//	for i := 0; i < size; i++ {
//		root[i] = i
//		rank[i] = 0
//	}
//	return UnionFound{root: root, rank: rank}
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
//func (uf *UnionFound) Union(x, y int) bool {
//	rootX, rootY := uf.Find(x), uf.Find(y)
//	if rootY == rootX {
//		return false
//	}
//	if uf.rank[rootX] > uf.rank[rootY] {
//		uf.root[rootY] = rootX
//	} else if uf.rank[rootY] > uf.rank[rootX] {
//		uf.root[rootX] = rootY
//	} else {
//		uf.rank[rootX]++
//		uf.root[rootY] = rootX
//	}
//	return true
//}

func minCostConnectPoints(points [][]int) int {
	visited := make([]bool, len(points))
	current, result, resultCount := 0, 0, 0
	distance := map[int][]*Edge{} // from, to, dist
	for i := 0; i < len(points); i++ {
		for k := i + 1; k < len(points); k++ {
			d := dist(points[i][0], points[k][0]) + dist(points[i][1], points[k][1])
			distance[i] = append(distance[i], &Edge{k, d})
			distance[k] = append(distance[k], &Edge{i, d})
		}
	}
	h := &MinHeap{}
	heap.Init(h)
	for _, v := range distance[current] {
		heap.Push(h, v)
	}
	visited[current] = true

	for h.Len() > 0 && resultCount < len(points) {
		candidate := heap.Pop(h).(*Edge)
		if !visited[candidate.to] {
			resultCount++
			visited[candidate.to] = true
			current = candidate.to
			result += candidate.cost
			for _, v := range distance[current] {
				if !visited[v.to] {
					heap.Push(h, v)
				}
			}
		}
	}
	return result
}

func dist(x, y int) int {
	result := x - y
	if result > 0 {
		return result
	} else {
		return -result
	}
}

type Edge struct {
	to   int
	cost int
}
type MinHeap []*Edge // to, distance

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(element interface{}) {
	*h = append(*h, element.(*Edge))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return element
}

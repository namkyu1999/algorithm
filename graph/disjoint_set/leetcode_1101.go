package main

//func main() {
//	n := 6
//	logs := [][]int{
//		{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5},
//	}
//	result := earliestAcq(logs, n)
//	fmt.Print(result)
//}
//
//type UnionFound struct {
//	rank []int
//	root []int
//}
//
//func Construct(size int) UnionFound {
//	rank, root := make([]int, size), make([]int, size)
//	for i := 0; i < size; i++ {
//		rank[i] = 1
//		root[i] = i
//	}
//	return UnionFound{
//		rank: rank, root: root,
//	}
//}
//func (uf *UnionFound) Find(x int) int {
//	if x == uf.root[x] {
//		return x
//	}
//	uf.root[x] = uf.Find(uf.root[x])
//	return uf.root[x]
//}
//
//func (uf *UnionFound) UnionWithValidation(x, y int) {
//	rootX, rootY := uf.Find(x), uf.Find(y)
//	if rootX != rootY {
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
//
//func earliestAcq(logs [][]int, n int) int {
//	uf := Construct(n)
//	sort.Slice(logs, func(i, j int) bool {
//		return logs[i][0] < logs[j][0]
//	})
//	for _, v := range logs {
//		uf.UnionWithValidation(v[1], v[2])
//		count := 0
//		for i, v := range uf.root {
//			if i == v {
//				count++
//			}
//		}
//		if count == 1 {
//			return v[0]
//		}
//	}
//	return -1
//}

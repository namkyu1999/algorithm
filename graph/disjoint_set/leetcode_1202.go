package main

//func main() {
//	s, pairs := "dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}
//	result := smallestStringWithSwaps(s, pairs)
//	fmt.Println(result)
//}

//func smallestStringWithSwaps(s string, pairs [][]int) string {
//	uf := InitUnionFound(len(s))
//	for _, v := range pairs {
//		uf.Union(v[0], v[1])
//	}
//	count := 0
//	for i, v := range uf.root {
//		if i == v {
//			count++
//		}
//	}
//	sRune := []rune(s)
//	if count == 1 {
//		sort.Slice(sRune, func(i, j int) bool {
//			return sRune[i] < sRune[j]
//		})
//		return string(sRune)
//	} else {
//		m := map[int][]rune{}
//		for i, v := range sRune {
//			m[uf.Find(i)] = append(m[uf.Find(i)], v)
//		}
//		for _, v := range m {
//			sort.Slice(v, func(i, j int) bool {
//				return v[i] < v[j]
//			})
//		}
//		result := make([]rune, len(sRune))
//		for i := 0; i < len(sRune); i++ {
//			if v, ok := m[uf.Find(i)]; ok {
//				result[i] = v[0]
//				m[uf.Find(i)] = v[1:]
//			} else {
//				result[i] = sRune[i]
//			}
//		}
//		return string(result)
//	}
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

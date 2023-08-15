package main

//import (
//	"fmt"
//)
//
//func main() {
//	equations := [][]string{{"a", "b"}, {"b", "c"}}
//	values := []float64{2.0, 3.0}
//	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
//	result := calcEquation(equations, values, queries)
//	fmt.Println(result)
//}
//
//func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
//	uf := UnionFound{
//		root:  map[string]string{},
//		rank:  map[string]int{},
//		value: map[string]float64{},
//	}
//	for i, v := range equations {
//		uf.Union(v[0], v[1], values[i])
//	}
//	result := make([]float64, len(queries))
//	for i, v := range queries {
//		if root0, ok := uf.root[v[0]]; !ok {
//			result[i] = -1
//			continue
//		} else if root1, ok := uf.root[v[1]]; !ok {
//			result[i] = -1
//			continue
//		} else {
//			if root1 != root0 {
//				result[i] = -1
//				continue
//			}
//			result[i] = uf.value[v[0]] / uf.value[v[1]]
//		}
//	}
//	return result
//}
//
//type UnionFound struct {
//	root  map[string]string
//	value map[string]float64
//	rank  map[string]int
//}
//
//func (uf *UnionFound) Find(x string) string {
//	return uf.root[x]
//}
//
//func (uf *UnionFound) Union(x, y string, value float64) {
//	if _, ok := uf.root[x]; !ok {
//		uf.root[x] = x
//		uf.value[x] = 1.0
//		uf.rank[x] = 1
//	}
//	if _, ok := uf.root[y]; !ok {
//		uf.root[y] = y
//		uf.value[y] = 1.0
//		uf.rank[y] = 1
//	}
//	rootX, rootY := uf.Find(x), uf.Find(y)
//	if rootX == rootY {
//		return
//	}
//	if uf.rank[rootX] > uf.rank[rootY] {
//		old := uf.value[y]
//		uf.value[y] = uf.value[x] / value
//		for k, v := range uf.root {
//			if v == rootY {
//				uf.root[k] = rootX
//				if k == y {
//					continue
//				}
//				uf.value[k] = uf.value[k] * uf.value[y] / old
//			}
//		}
//	} else if uf.rank[rootX] < uf.rank[rootY] {
//		old := uf.value[x]
//		uf.value[x] = uf.value[y] * value
//		for k, v := range uf.root {
//			if v == rootX {
//				uf.root[k] = rootY
//				if k == x {
//					continue
//				}
//				uf.value[k] = uf.value[k] * uf.value[x] / old
//			}
//		}
//	} else {
//		uf.rank[rootX]++
//		old := uf.value[y]
//		uf.value[y] = uf.value[x] / value
//		for k, v := range uf.root {
//			if v == rootY {
//				uf.root[k] = rootX
//				if k == y {
//					continue
//				}
//				uf.value[k] = uf.value[k] * uf.value[y] / old
//			}
//		}
//	}
//}

//package main
//
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
//		root0, exist0 := uf.Find(v[0])
//		root1, exist1 := uf.Find(v[1])
//		if exist1 && exist0 && root1 == root0 {
//			result[i] = uf.value[v[0]] / uf.value[v[1]]
//		} else {
//			result[i] = -1
//		}
//	}
//	return result
//}
//
//type UnionFound struct {
//	rank  map[string]int
//	root  map[string]string
//	value map[string]float64
//}
//
//func (uf *UnionFound) Find(x string) (string, bool) {
//	if v, ok := uf.root[x]; !ok {
//		return "", false
//	} else {
//		if x == v {
//			return x, true
//		} else {
//			newOne, _ := uf.Find(v)
//			uf.root[x] = newOne
//			return uf.root[x], true
//		}
//	}
//}
//
//func (uf *UnionFound) Union(x, y string, value float64) {
//	if _, ok := uf.value[x]; ok {
//		uf.value[y] = uf.value[x] / value
//	} else if _, ok := uf.value[x]; ok {
//		uf.value[x] = uf.value[y] * value
//	} else {
//		uf.value[x] = value
//		uf.value[y] = 1
//	}
//	if _, ok := uf.root[x]; !ok {
//		uf.root[x] = x
//	}
//	if _, ok := uf.root[y]; !ok {
//		uf.root[y] = y
//	}
//	rootX, _ := uf.Find(x)
//	rootY, _ := uf.Find(y)
//	if rootX != rootY {
//		if _, ok := uf.rank[x]; !ok {
//			uf.rank[x] = 0
//		}
//		if _, ok := uf.rank[y]; !ok {
//			uf.rank[y] = 0
//		}
//		if uf.rank[rootX] > uf.rank[rootY] {
//			uf.root[rootY] = rootX
//		} else if uf.rank[rootY] > uf.rank[rootX] {
//			uf.root[rootX] = rootY
//		} else {
//			uf.rank[rootX]++
//			uf.root[rootY] = rootX
//		}
//	}
//}

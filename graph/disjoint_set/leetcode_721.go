package main

//func main() {
//	accounts := [][]string{
//		{"David", "David0@m.co", "David1@m.co"},
//		{"David", "David3@m.co", "David4@m.co"},
//		{"David", "David4@m.co", "David5@m.co"},
//		{"David", "David2@m.co", "David3@m.co"},
//		{"David", "David1@m.co", "David2@m.co"},
//	}
//	result := accountsMerge(accounts)
//	fmt.Println(result)
//}
//
//func accountsMerge(accounts [][]string) [][]string {
//	uf := InitUnionFound()
//	nameList := map[string]string{}
//
//	for _, v := range accounts {
//		if _, ok := uf.root[v[1]]; !ok {
//			uf.root[v[1]] = v[1]
//		}
//		nameList[v[1]] = v[0]
//		for _, a := range v[2:] {
//			uf.Union(v[1], a)
//			nameList[a] = v[0]
//		}
//	}
//	result := map[string][]string{}
//	for k := range uf.root {
//		result[uf.Find(k)] = append(result[uf.Find(k)], k)
//	}
//	var answer [][]string
//	for k, v := range result {
//		sort.Strings(v)
//		answer = append(answer, append([]string{nameList[k]}, v...))
//	}
//	return answer
//}
//
//type UnionFound struct {
//	rank map[string]int
//	root map[string]string
//}
//
//func InitUnionFound() UnionFound {
//	return UnionFound{
//		root: map[string]string{},
//		rank: map[string]int{},
//	}
//}
//
//func (uf *UnionFound) Find(x string) string {
//	if x == uf.root[x] {
//		return x
//	}
//	uf.root[x] = uf.Find(uf.root[x])
//	return uf.root[x]
//}
//
//func (uf *UnionFound) Union(x, y string) {
//	if _, ok := uf.root[x]; !ok {
//		uf.root[x] = x
//	}
//	if _, ok := uf.root[y]; !ok {
//		uf.root[y] = y
//	}
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

# Disjoint Set
> leetcode learn card - graph
- The find function finds the root node of a given vertex.
- The union function unions two vertices and makes their root nodes the same.
> The main idea of a “disjoint set” is to have all connected vertices have the same parent node or root node, whether directly or indirectly connected. To check if two vertices are connected, we only need to check if they have the same root node.
## Quick Find
```go
// Quick Find
func main() {
	uf := InitUnionFind(10)
	// 1-2-5-6-7 3-8-9 4
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	fmt.Printf("connected 1,5: %t\n", uf.Connected(1, 5)) // true
	fmt.Printf("connected 5,7: %t\n", uf.Connected(5, 7)) // true
	fmt.Printf("connected 4,9: %t\n", uf.Connected(4, 9)) // false
	// 1-2-5-6-7 3-8-9-4
	uf.Union(9, 4)
	fmt.Printf("connected 9,4: %t\n", uf.Connected(4, 9)) // true
}

type UnionFind struct {
	root []int
}

func InitUnionFind(size int) UnionFind {
	uf := UnionFind{}
	uf.root = make([]int, size)
	for i := 0; i < size; i++ {
		uf.root[i] = i
	}
	return uf
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		for i := 0; i < len(uf.root); i++ {
			if uf.root[i] == rootY {
				uf.root[i] = rootX
			}
		}
	}
}

func (uf *UnionFind) Find(x int) int {
	return uf.root[x]
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
```
## Quick Union
```go
// Quick Union
func main() {
	uf := InitUnionFind(10)
	// 1-2-5-6-7 3-8-9 4
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	fmt.Printf("connected 1,5: %t\n", uf.Connected(1, 5)) // true
	fmt.Printf("connected 5,7: %t\n", uf.Connected(5, 7)) // true
	fmt.Printf("connected 4,9: %t\n", uf.Connected(4, 9)) // false
	// 1-2-5-6-7 3-8-9-4
	uf.Union(9, 4)
	fmt.Printf("connected 9,4: %t\n", uf.Connected(4, 9)) // true
}

type UnionFind struct {
	root []int
}

func InitUnionFind(size int) UnionFind {
	uf := UnionFind{}
	uf.root = make([]int, size)
	for i := 0; i < size; i++ {
		uf.root[i] = i
	}
	return uf
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		uf.root[rootY] = rootX
	}
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.root[x] {
		x = uf.root[x]
	}
	return x
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
```
## Union by Rank
```go
func main() {
	uf := InitUnionFind(10)
	// 1-2-5-6-7 3-8-9 4
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	fmt.Printf("connected 1,5: %t\n", uf.Connected(1, 5)) // true
	fmt.Printf("connected 5,7: %t\n", uf.Connected(5, 7)) // true
	fmt.Printf("connected 4,9: %t\n", uf.Connected(4, 9)) // false
	// 1-2-5-6-7 3-8-9-4
	uf.Union(9, 4)
	fmt.Printf("connected 9,4: %t\n", uf.Connected(4, 9)) // true
}

type UnionFind struct {
	root []int
	rank []int
}

func InitUnionFind(size int) UnionFind {
	uf := UnionFind{}
	uf.root, uf.rank = make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = uf.root[rootX]
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootX] = uf.root[rootY]
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.root[x] {
		x = uf.root[x]
	}
	return x
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
```
## Path Compression
```go
func (uf *UnionFind) Find(x int) int {
	if x == uf.root[x] {
		return x
	}
	uf.root[x] = uf.Find(uf.root[x])
	return uf.root[x]
}
```

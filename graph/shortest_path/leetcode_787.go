package main

//func main() {
//	n, src, dst, k := 4, 0, 3, 1
//	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
//	result := findCheapestPrice(n, flights, src, dst, k)
//	fmt.Println(result)
//}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prev, cur := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		prev[i], cur[i] = 100001, 100001
	}
	prev[src], cur[src] = 0, 0
	for i := 0; i <= k; i++ {
		for _, v := range flights {
			candidate := prev[v[0]] + v[2]
			if cur[v[1]] > candidate {
				cur[v[1]] = candidate
			}
		}
		copy(prev, cur)
	}
	if cur[dst] == 100001 {
		return -1
	}
	return cur[dst]
}

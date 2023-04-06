package main

import "container/heap"

//func main() {
//	n, src, dst, k := 4, 0, 3, 1
//	flights := [][]int{
//		{0, 1, 100},
//		{1, 2, 100},
//		{2, 0, 100},
//		{1, 3, 600},
//		{2, 3, 200},
//	}
//	result := findCheapestPrice(n, flights, src, dst, k)
//	fmt.Println(result)
//}

type FlightHeap []Flight

func (h FlightHeap) Len() int { return len(h) }

func (h FlightHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h FlightHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *FlightHeap) Push(x interface{}) {
	*h = append(*h, x.(Flight))
}

func (h *FlightHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func prepareFlights(edges [][]int, n int) ([][]int, [][]int) {
	neighbors := make([][]int, n)
	costs := make([][]int, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make([]int, 0)
		costs[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		k, v, cost := edges[i][0], edges[i][1], edges[i][2]
		neighbors[k] = append(neighbors[k], v)
		costs[k] = append(costs[k], cost)
	}
	return neighbors, costs
}

type Flight struct {
	pos   int
	cost  int
	stops int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	neighbors, costs := prepareFlights(flights, n)
	stops := make([]int, n)
	for i := 0; i < n; i++ {
		stops[i] = 100000000000
	}

	current := Flight{
		pos:   src,
		cost:  0,
		stops: 0,
	}
	h := &FlightHeap{current}
	heap.Init(h)
	for h.Len() > 0 {
		node := heap.Pop(h).(Flight)
		if node.stops > stops[node.pos] || node.stops > k+1 {
			continue
		}
		stops[node.pos] = node.stops
		if node.pos == dst {
			return node.cost
		}
		for i, nei := range neighbors[node.pos] {
			heap.Push(h, Flight{
				pos:   nei,
				cost:  node.cost + costs[node.pos][i],
				stops: node.stops + 1,
			})
		}
	}
	return -1
}

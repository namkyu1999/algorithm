package main

import "fmt"

//func main() {
//	n := 5
//	redEdges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
//	blueEdges := [][]int{{1, 2}, {2, 3}, {3, 1}}
//	result := shortestAlternatingPaths(n, redEdges, blueEdges)
//	fmt.Println(result)
//}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type Node struct{ nodeIdx, distance, color int }

	res := make([]int, n)
	for i := 1; i < n; i++ {
		res[i] = -1
	}
	adjRed := map[int][]int{}
	adjBlue := map[int][]int{}

	for _, edge := range redEdges {
		adjRed[edge[0]] = append(adjRed[edge[0]], edge[1])
	}

	for _, edge := range blueEdges {
		adjBlue[edge[0]] = append(adjBlue[edge[0]], edge[1])
	}

	// color 0 - special case for root node
	q := []Node{{nodeIdx: 0, distance: 0, color: 0}}

	// colors 1: red, 2: blue
	const red, blue = 1, 2

	visited := map[string]bool{}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if res[cur.nodeIdx] == -1 {
			res[cur.nodeIdx] = cur.distance
		}

		// if not red
		if cur.color != red {
			for _, nei := range adjRed[cur.nodeIdx] {
				v := fmt.Sprintf("%d->%d", nei, red)
				if visited[v] {
					continue
				}
				visited[v] = true
				q = append(q, Node{nodeIdx: nei, distance: cur.distance + 1, color: red})
			}

		}

		// if not blue
		if cur.color != blue {
			for _, nei := range adjBlue[cur.nodeIdx] {
				v := fmt.Sprintf("%d->%d", nei, blue)
				if visited[v] {
					continue
				}
				visited[v] = true
				q = append(q, Node{nodeIdx: nei, distance: cur.distance + 1, color: blue})
			}
		}

	}

	return res
}

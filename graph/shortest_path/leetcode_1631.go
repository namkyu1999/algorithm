package main

import (
	"math"
)

//func main() {
//	heights := [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
//	result := minimumEffortPath(heights)
//	fmt.Println(result)
//}

func minimumEffortPath(heights [][]int) int {
	maxRow, maxCol := len(heights), len(heights[0])
	outOfIndex := func(x, y int) bool {
		if x < 0 || y < 0 || x >= maxRow || y >= maxCol {
			return true
		}
		return false
	}
	efforts := make([][]int, maxRow)
	for i := 0; i < maxRow; i++ {
		efforts[i] = make([]int, maxCol)
		for k := 0; k < maxCol; k++ {
			efforts[i][k] = math.MaxInt
		}
	}
	efforts[0][0] = 0
	queue := [][2]int{{0, 0}}
	direction := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, pos := range direction {
			nextX, nextY := current[0]+pos[0], current[1]+pos[1]
			if outOfIndex(nextX, nextY) {
				continue
			}
			candidate := heights[nextX][nextY] - heights[current[0]][current[1]]
			if candidate < 0 {
				candidate = -candidate
			}
			if candidate < efforts[current[0]][current[1]] {
				candidate = efforts[current[0]][current[1]]
			}
			if efforts[nextX][nextY] > candidate {
				efforts[nextX][nextY] = candidate
				queue = append(queue, [2]int{nextX, nextY})
			}
		}
	}
	return efforts[maxRow-1][maxCol-1]
}

package main

//func main() {
//	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
//	result := shortestPathBinaryMatrix(grid)
//	fmt.Println(result)
//}

func shortestPathBinaryMatrix(grid [][]int) int {
	lenX, lenY := len(grid), len(grid[0])
	canNext := func(x, y int) bool {
		if x > -1 && y > -1 && x < lenX && y < lenY && grid[x][y] == 0 {
			return true
		}
		return false
	}
	direction := [][2]int{{1, 1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, -1}, {-1, 1}, {-1, -1}}
	if len(grid) == 0 || grid[0][0] == 1 {
		return -1
	}
	queue := [][3]int{{0, 0, 1}}
	for len(queue) > 0 {
		x, y, step := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if x == lenX-1 && y == lenY-1 {
			return step
		}
		for _, d := range direction {
			if canNext(x+d[0], y+d[1]) {
				grid[x+d[0]][y+d[1]] = 1
				queue = append(queue, [3]int{x + d[0], y + d[1], step + 1})
			}
		}
	}
	return -1
}

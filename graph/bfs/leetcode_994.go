package main

//func main() {
//	// 2: rotten, 1: fresh, 0: none
//	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
//	result := orangesRotting(grid)
//	fmt.Println(result)
//}

func orangesRotting(grid [][]int) int {
	days := -1
	var queue [][2]int
	noFresh := true
	direction := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	lenX, lenY := len(grid), len(grid[0])
	canNext := func(x, y int) bool {
		if x > -1 && x < lenX && y > -1 && y < lenY && grid[x][y] == 1 {
			return true
		}
		return false
	}
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if grid[x][y] == 2 {
				queue = append(queue, [2]int{x, y})
			} else if grid[x][y] == 1 {
				noFresh = false
			}
		}
	}
	if noFresh {
		return 0
	}
	for len(queue) > 0 {
		days++
		length := len(queue)
		for i := 0; i < length; i++ {
			current := queue[i]
			x, y := current[0], current[1]
			for _, d := range direction {
				if canNext(x+d[0], y+d[1]) {
					grid[x+d[0]][y+d[1]] = 2
					queue = append(queue, [2]int{x + d[0], y + d[1]})
				}
			}
		}
		queue = queue[length:]
	}
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if grid[x][y] == 1 {
				return -1
			}
		}
	}
	return days
}

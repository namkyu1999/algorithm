package main

import (
	"math"
)

//func main() {
//	maze := [][]byte{
//		{'+', '+', '+'},
//		{'.', '.', '.'},
//		{'+', '+', '+'},
//	}
//	entrance := []int{1, 0}
//	result := nearestExit(maze, entrance)
//	fmt.Println(result)
//}

func nearestExit(maze [][]byte, entrance []int) int {
	notInMaze := func(x, y int) bool {
		if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) {
			return true
		}
		return false
	}
	result := math.MaxInt
	direction := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	maze[entrance[0]][entrance[1]] = '+'
	type status struct {
		step    int
		current [2]int
	}
	queue := []status{{step: 0, current: [2]int{entrance[0], entrance[1]}}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			next := [2]int{pos.current[0] + v[0], pos.current[1] + v[1]}
			if notInMaze(next[0], next[1]) || maze[next[0]][next[1]] == '+' {
				continue
			} else if (next[0] == 0 || next[1] == 0 || next[0] == len(maze)-1 || next[1] == len(maze[0])-1) && maze[next[0]][next[1]] != '+' {
				if result > pos.step+1 {
					result = pos.step + 1
				}
			} else {
				maze[next[0]][next[1]] = '+'
				queue = append(queue, status{
					step:    pos.step + 1,
					current: next,
				})
			}
		}
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}

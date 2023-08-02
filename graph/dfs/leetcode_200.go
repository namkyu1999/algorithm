package main

//func main() {
//	input := [][]byte{
//		[]byte{byte('1'), byte('1'), byte('1'), byte('1'), byte('0')},
//		[]byte{byte('1'), byte('1'), byte('0'), byte('1'), byte('0')},
//		[]byte{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
//		[]byte{byte('0'), byte('0'), byte('0'), byte('1'), byte('0')},
//	}
//	result := numIslands(input)
//	fmt.Println(result)
//}

func numIslands(grid [][]byte) int {
	result := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == byte('1') {
				ndfs(x, y, &grid)
				result++
			}
		}
	}
	return result
}

func ndfs(x, y int, grid *[][]byte) {
	if (*grid)[x][y] == byte('0') {
		return
	}
	(*grid)[x][y] = byte('0')
	forward := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, f := range forward {
		if x+f[0] < len(*grid) && -1 < x+f[0] && y+f[1] < len((*grid)[0]) && -1 < y+f[1] {
			ndfs(x+f[0], y+f[1], grid)
		}
	}
}

//// ndfs
//func numIslands(grid [][]byte) int {
//	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
//	m, n := len(grid), len(grid[0])
//	result := 0
//	for i := 0; i < m; i++ {
//		for k := 0; k < n; k++ {
//			if grid[i][k] == byte('1') {
//				grid[i][k] = byte('0')
//				for _, offset := range offsets {
//					row, col := i+offset[0], k+offset[1]
//					if 0 <= row && row < m && 0 <= col && col < n {
//						ndfs(row, col, grid)
//					}
//				}
//				result++
//			}
//		}
//	}
//	return result
//}

//func ndfs(row int, col int, grid [][]byte) {
//	if grid[row][col] == byte('1') {
//		offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
//		grid[row][col] = byte('0')
//		for _, offset := range offsets {
//			newRow, newCol := row+offset[0], col+offset[1]
//			if 0 <= newRow && newRow < len(grid) && 0 <= newCol && newCol < len(grid[0]) {
//				ndfs(newRow, newCol, grid)
//			}
//		}
//	}
//}
//
//// bfs
//func numIslands1(grid [][]byte) int {
//	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
//	m, n := len(grid), len(grid[0])
//	result := 0
//	for i := 0; i < m; i++ {
//		for k := 0; k < n; k++ {
//			if grid[i][k] == byte('1') {
//				result++
//				grid[i][k] = byte('0')
//				queue := make([][2]int, 0)
//				if safeIndex(m, n, i+1, k) {
//					queue = append(queue, [2]int{i + 1, k})
//				}
//				if safeIndex(m, n, i, k+1) {
//					queue = append(queue, [2]int{i, k + 1})
//				}
//				for len(queue) > 0 {
//					coor := queue[0]
//					queue = queue[1:]
//					if grid[coor[0]][coor[1]] == byte('1') {
//						grid[coor[0]][coor[1]] = byte('0')
//						for _, offset := range offsets {
//							if safeIndex(m, n, coor[0]+offset[0], coor[1]+offset[1]) {
//								queue = append(queue, [2]int{coor[0] + offset[0], coor[1] + offset[1]})
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//	return result
//}
//
//func safeIndex(m, n, i, k int) bool {
//	if i < m && k < n && k > -1 && i > -1 {
//		return true
//	}
//	return false
//}

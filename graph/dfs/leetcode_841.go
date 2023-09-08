package main

//func main() {
//	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
//	result := canVisitAllRooms(rooms)
//	fmt.Println(result)
//}

func canVisitAllRooms(rooms [][]int) bool {
	left := len(rooms)
	visited := make([]bool, left)

	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		left--
		for _, v := range rooms[i] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(0)
	if left != 0 {
		return false
	}
	return true
}

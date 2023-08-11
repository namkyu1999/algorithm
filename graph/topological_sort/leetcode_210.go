package main

//func main() {
//	numCourses := 4
//	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
//	result := findOrder(numCourses, prerequisites)
//	fmt.Println(result)
//}

func findOrder(numCourses int, prerequisites [][]int) []int {
	var result []int
	hash := map[int][]int{}
	inOrder, visited := make([]int, numCourses), make([]bool, numCourses)
	for _, v := range prerequisites {
		inOrder[v[0]]++
		hash[v[1]] = append(hash[v[1]], v[0])
	}
	var queue []int
	for i, v := range inOrder {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		result = append(result, current)
		for _, v := range hash[current] {
			inOrder[v]--
		}
		for i, v := range inOrder {
			if v == 0 && !visited[i] {
				queue = append(queue, i)
			}
		}
	}
	if len(result) < numCourses {
		return []int{}
	}
	return result
}

package main

//func main() {
//	n, relations := 3, [][]int{{1, 3}, {2, 3}}
//	result := minimumSemesters(n, relations)
//	fmt.Println(result)
//}

func minimumSemesters(n int, relations [][]int) int {
	visitedCount, result := 0, 0
	inOrder := make([]int, n+1)
	hash := map[int][]int{}
	for _, v := range relations {
		inOrder[v[1]]++
		hash[v[0]] = append(hash[v[0]], v[1])
	}
	var queue []int
	for i := 1; i <= n; i++ {
		if inOrder[i] == 0 {
			queue = append(queue, i)
			inOrder[i]--
		}
	}
	for len(queue) > 0 {
		for _, q := range queue {
			visitedCount++
			for _, v := range hash[q] {
				inOrder[v]--
			}
		}
		result++
		queue = []int{}
		for i := 1; i <= n; i++ {
			if inOrder[i] == 0 {
				queue = append(queue, i)
				inOrder[i]--
			}
		}
	}
	if visitedCount == n {
		return result
	}
	return -1
}

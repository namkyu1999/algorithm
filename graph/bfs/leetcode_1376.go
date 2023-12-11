package main

//func main() {
//	n := 15
//	headID := 0
//	manager := []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
//	informTime := []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
//	result := numOfMinutes(n, headID, manager, informTime)
//	fmt.Println(result)
//}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	type Node struct {
		index int
		time  int
	}
	hash := map[int][]int{}
	queue := []Node{{index: headID}}
	result := 0
	for i, v := range manager {
		hash[v] = append(hash[v], i)
	}
	for len(queue) > 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			current := queue[i]
			if len(hash[current.index]) == 0 {
				if current.time > result {
					result = current.time
				}
			}
			for _, v := range hash[current.index] {
				queue = append(queue, Node{index: v, time: current.time + informTime[current.index]})
			}
		}
		queue = queue[queueLength:]
	}
	return result
}

package main

import (
	"container/heap"
)

//func main() {
//	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
//	k := 4
//	result := findKthLargest(nums, k)
//	fmt.Println(result)
//}

func findKthLargest(nums []int, k int) int {
	maxHeap := MaxHeap(nums)
	heap.Init(&maxHeap)
	for i := 0; i < k-1; i++ {
		heap.Pop(&maxHeap)
	}
	return heap.Pop(&maxHeap).(int)
}

type MaxHeap []int

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	result := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return result
}

func (m MaxHeap) Len() int {
	return len(m)
}

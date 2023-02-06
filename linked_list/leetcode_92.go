package main

//func main() {
//	list, left, right := []int{1, 2, 3, 4, 5}, 2, 4
//	head := listToLinkedList(list)
//	result := reverseBetween(head, left, right)
//	printLinkedList(result)
//}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	root := &ListNode{Next: head}
	prev := root
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	pivot := prev
	prev = prev.Next
	current := prev.Next
	for left < right {
		l := pivot.Next
		pivot.Next = current
		prev.Next = current.Next
		current.Next = l
		current = prev.Next
		left++
	}
	return root.Next
}

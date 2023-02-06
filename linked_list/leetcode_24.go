package main

//func main() {
//	list := []int{1, 2, 3, 4}
//	head := listToLinkedList(list)
//	result := swapPairs(head)
//	fmt.Println("--")
//	printLinkedList(result)
//}

func swapPairs(head *ListNode) *ListNode {
	root := &ListNode{}
	prev := root
	prev.Next = head
	for head != nil && head.Next != nil {
		b := head.Next
		head.Next = b.Next
		b.Next = head
		prev.Next = b
		prev = prev.Next.Next
		head = prev.Next
	}
	return root.Next
}

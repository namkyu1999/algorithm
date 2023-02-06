package main

//func main() {
//	list := []int{1, 2, 3, 4, 5}
//	head := listToLinkedList(list)
//	result := oddEvenList(head)
//	printLinkedList(result)
//}

func oddEvenList(head *ListNode) *ListNode {
	root := &ListNode{}
	prev := root
	prev.Next = head
	evenRoot := &ListNode{}
	evenHead := evenRoot
	isOdd := true
	for head != nil {
		if !isOdd {
			prev.Next = head.Next
			b := head
			b.Next = nil
			evenHead.Next = b
			evenHead = evenHead.Next
			head = prev.Next
		} else {
			prev = prev.Next
			head = head.Next
		}
		isOdd = !isOdd
	}
	prev.Next = evenRoot.Next
	return root.Next
}

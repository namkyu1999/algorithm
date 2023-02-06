package main

//func main() {
//	list1, list2 := []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}
//	l1, l2 := listToLinkedList(list1), listToLinkedList(list2)
//	answer := addTwoNumbers(l1, l2)
//	printLinkedList(answer)
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	result := &ListNode{Val: (l1.Val + l2.Val) % 10}
	current := result
	add := (l1.Val + l2.Val) / 10
	l1, l2 = l1.Next, l2.Next
	for l1 != nil || l2 != nil {
		if l1 == nil {
			current.Next = &ListNode{Val: (l2.Val + add) % 10}
			current = current.Next
			add = (l2.Val + add) / 10
			l2 = l2.Next
			continue
		} else if l2 == nil {
			current.Next = &ListNode{Val: (l1.Val + add) % 10}
			current = current.Next
			add = (l1.Val + add) / 10
			l1 = l1.Next
			continue
		}
		current.Next = &ListNode{Val: (l1.Val + l2.Val + add) % 10}
		add = (l1.Val + l2.Val + add) / 10
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if add == 1 {
		current.Next = &ListNode{Val: 1}
	}
	return result
}

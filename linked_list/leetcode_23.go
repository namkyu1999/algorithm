package main

//func main() {
//	lists := []*ListNode{
//		listToLinkedList([]int{1, 3, 4, 6, 8, 9, 12}),
//		listToLinkedList([]int{1, 2, 5, 7, 11, 21, 24}),
//		listToLinkedList([]int{-4, 0, 4, 7, 10, 14, 22, 29}),
//	}
//	answer := mergeKLists(lists)
//	printLinkedList(answer)
//}

func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	default:
		root := &ListNode{}
		prev := root
		prev.Next = lists[0]
		current := prev.Next
		for _, list := range lists[1:] {
			if list == nil {
				continue
			}
			if current == nil {
				prev.Next = list
				current = prev.Next
				continue
			}
			for list != nil {
				if current.Val < list.Val {
					prev = prev.Next
					current = current.Next
				} else {
					l := list
					list = list.Next
					prev.Next = l
					l.Next = current
					current = prev.Next
				}
				if current == nil {
					prev.Next = list
					break
				}
			}
			prev = root
			current = prev.Next
		}
		return root.Next
	}
}

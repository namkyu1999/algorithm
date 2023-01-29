package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func main() {
//	list1, list2 := []int{1, 2, 4}, []int{1, 3, 4}
//	linkedList1, linkedList2 := listToLinkedList(list1), listToLinkedList(list2)
//	result := mergeTwoLists(linkedList1, linkedList2)
//	printLinkedList(result)
//}

func printLinkedList(list1 *ListNode) {
	current := list1
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func listToLinkedList(list []int) *ListNode {
	listSize := len(list)
	switch listSize {
	case 0:
		return nil
	case 1:
		return &ListNode{Val: list[0]}
	default:
		result := &ListNode{Val: list[0]}
		current := result
		for i := 1; i < listSize; i++ {
			current.Next = &ListNode{Val: list[i]}
			current = current.Next
		}
		return result
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var result *ListNode
	if list1.Val > list2.Val {
		result = &ListNode{Val: list2.Val}
		list2 = list2.Next
	} else {
		result = &ListNode{Val: list1.Val}
		list1 = list1.Next
	}
	current := result

	for list1 != nil || list2 != nil {
		if list1 == nil {
			current.Next = list2
			break
		} else if list2 == nil {
			current.Next = list1
			break
		}
		if list1.Val > list2.Val {
			current.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			current.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		current = current.Next
	}

	return result
}

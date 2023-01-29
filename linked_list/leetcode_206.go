package main

//func main() {
//	input := []int{1, 2, 3, 4, 5}
//	linkedList := listToLinkedList(input)
//	answer := reverseList(linkedList)
//	printLinkedList(answer)
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	result := reverseList(head.Next)
	current := result
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: head.Val}
	return result
}

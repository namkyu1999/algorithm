package main

import "fmt"

//func main() {
//	raw := []int{-1, 5, 8, 11, 12}
//	prev := &ListNode{}
//	current := prev
//	for _, v := range raw {
//		current.Next = &ListNode{Val: v}
//		current = current.Next
//	}
//
//	result := sortList(prev.Next)
//	printLinkedList(result)
//}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)
	return assembleNode(left, right)
}

func assembleNode(head1, head2 *ListNode) *ListNode {
	curHead := &ListNode{}
	tmpHead := curHead

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			curHead.Next = head1
			head1 = head1.Next
			curHead = curHead.Next
		} else {
			curHead.Next = head2
			head2 = head2.Next
			curHead = curHead.Next
		}
	}
	if head1 != nil {
		curHead.Next = head1
	} else if head2 != nil {
		curHead.Next = head2
	}
	return tmpHead.Next
}

package main

//func main() {
//	input := []int{1, 2, 2, 1}
//	head := &ListNode{Val: input[0]}
//	current := head
//	for _, v := range input[1:] {
//		current.Next = &ListNode{Val: v}
//		current = current.Next
//	}
//	result := isPalindrome(head)
//	fmt.Println(result)
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	current := head
	arr := make([]int, 0)
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

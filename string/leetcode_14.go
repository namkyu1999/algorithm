package main

func longestCommonPrefix(strs []string) string {
	root := &prefixNode{}
	current := root
	for _, v := range []rune(strs[0]) {
		current.Next = &prefixNode{Val: v}
		current = current.Next
	}
	for _, v := range strs[1:] {
		prev := root
		current = root.Next
		if current == nil {
			return ""
		}
		for _, word := range []rune(v) {
			if current.Val == word {
				prev = current
				current = current.Next
			} else {
				prev.Next = nil
				current = nil
			}
			if current == nil {
				break
			}
		}
		if current != nil {
			prev.Next = nil
		}
	}
	result := ""
	current = root.Next
	for current != nil {
		result += string(current.Val)
		current = current.Next
	}
	return result
}

type prefixNode struct {
	Val  rune
	Next *prefixNode
}

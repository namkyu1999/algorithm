package main

//
//import "fmt"
//
//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}
//
//func main() {
//	root := &Node{
//		Val:   1,
//		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
//		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}},
//	}
//	result := connect(root)
//	fmt.Println(result)
//}
//
//func connect(root *Node) *Node {
//	ratio := 1
//	prev, wave := 0, 0
//	queue := []*Node{root}
//	for len(queue) > 0 {
//		wave++
//		current := queue[0]
//		queue = queue[1:]
//		if current != nil && current.Left != nil {
//			queue = append(queue, current.Left, current.Right)
//		}
//		if prev+ratio == wave {
//			prev = wave
//			ratio *= 2
//			current.Next = nil
//			continue
//		}
//		current.Next = queue[0]
//	}
//	return root
//}

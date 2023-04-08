package main

//func main() {
//	root := &TreeNode{
//		Val: 3,
//		Left: &TreeNode{
//			Val: 9,
//		},
//		Right: &TreeNode{
//			Val: 20,
//			Left: &TreeNode{
//				Val: 15,
//			},
//			Right: &TreeNode{
//				Val: 7,
//			},
//		},
//	}
//	result := maxDepth(root)
//	fmt.Println(result)
//}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

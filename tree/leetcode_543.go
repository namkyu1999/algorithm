package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func main() {
//	root := &TreeNode{
//		Val: 1,
//		Left: &TreeNode{
//			Val: 2,
//		},
//		Right: &TreeNode{
//			Val: 3,
//			Left: &TreeNode{
//				Val: 15,
//			},
//			Right: &TreeNode{
//				Val: 7,
//			},
//		},
//	}
//	result := diameterOfBinaryTree(root)
//	fmt.Println(result)
//}

var longest = 0

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)

	if longest < left+right+2 {
		longest = left + right + 2
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

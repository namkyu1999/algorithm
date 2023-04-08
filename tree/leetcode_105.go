package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//func main() {
//	preorder := []int{3, 9, 20, 15, 7}
//	inorder := []int{9, 3, 15, 20, 7}
//	result := buildTree(preorder, inorder)
//	fmt.Println(result)
//}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)

	if n == 0 {
		return nil
	}

	pv := preorder[0]
	pi := 0
	for pi < n && inorder[pi] != pv {
		pi++
	}

	ans := new(TreeNode)
	ans.Val = pv
	ans.Left = buildTree(preorder[1:], inorder[:pi])
	ans.Right = buildTree(preorder[1+pi:], inorder[pi+1:])
	return ans
}

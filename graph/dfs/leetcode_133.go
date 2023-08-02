package main

//type Node struct {
//	Val       int
//	Neighbors []*Node
//}
//
//func cloneGraph(node *Node) *Node {
//	if node == nil {
//		return nil
//	}
//	nextMap := map[int]*Node{}
//	var dfs func(n *Node) *Node
//	dfs = func(n *Node) *Node {
//		newNode := new(Node)
//		newNode.Val = n.Val
//		nextMap[n.Val] = newNode
//		for _, v := range n.Neighbors {
//			if nextMap[v.Val] == nil {
//				nextMap[v.Val] = dfs(v)
//			}
//			newNode.Neighbors = append(newNode.Neighbors, nextMap[v.Val])
//		}
//		return newNode
//	}
//	return dfs(node)
//}

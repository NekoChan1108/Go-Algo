// 自己写的(二叉树中序遍历)
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	nodeSlice := make([]int, 0)
	if root == nil {
		return nodeSlice
	}
	nodeSlice = append(nodeSlice, inorderTraversal(root.Left)...)
	nodeSlice = append(nodeSlice, root.Val)
	nodeSlice = append(nodeSlice, inorderTraversal(root.Right)...)
	return nodeSlice
}

func main() {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Right = n4
	fmt.Println(inorderTraversal(n1))
}

// 自己写的(二叉树最大深度)
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	r1 := &TreeNode{1, nil, nil}
	r2 := &TreeNode{2, nil, nil}
	r3 := &TreeNode{3, nil, nil}
	r4 := &TreeNode{4, nil, nil}
	r5 := &TreeNode{5, nil, nil}
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r4.Left = r5
	fmt.Print(maxDepth(r1))
}

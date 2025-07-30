// 自己写的(先序遍历)
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nodeSlice = make([]int, 0)
	nodeSlice = append(nodeSlice, root.Val)
	nodeSlice = append(nodeSlice, preorderTraversal(root.Left)...)
	nodeSlice = append(nodeSlice, preorderTraversal(root.Right)...)
	return nodeSlice
}

func main() {
	r1 := &TreeNode{1, nil, nil}
	r2 := &TreeNode{2, nil, nil}
	r3 := &TreeNode{3, nil, nil}
	r4 := &TreeNode{4, nil, nil}
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	fmt.Println(preorderTraversal(r1))
}

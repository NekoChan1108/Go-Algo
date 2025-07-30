// 自己写的(最长直径)
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//存最大左右子树高度之和
var dis int

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var lh int = height(root.Left)
	var rh int = height(root.Right)
	dis = max(dis, lh+rh)
	// fmt.Println(dis)
	return max(lh, rh) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	height(root)
	return dis
}

func main() {
	r1 := &TreeNode{1, nil, nil}
	r2 := &TreeNode{2, nil, nil}
	// r3 := &TreeNode{3, nil, nil}
	// r4 := &TreeNode{4, nil, nil}
	// r5 := &TreeNode{5, nil, nil}
	// r6 := &TreeNode{6, nil, nil}

	r1.Left = r2
	// r1.Right = r4
	// r2.Left = r3
	// r3.Right = r5
	// r5.Right = r6

	fmt.Print(diameterOfBinaryTree(r1))
}

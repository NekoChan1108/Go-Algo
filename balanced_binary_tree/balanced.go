// 自己写的(判断是否是二叉平衡树)
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(height(root.Left), height(root.Right)) + 1
}

// 左子树与右子树高度之差的绝对值<=1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//根节点与依次递归的左节点和右节点全部满足
	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
	r1 := &TreeNode{1, nil, nil}
	r2 := &TreeNode{2, nil, nil}
	r3 := &TreeNode{3, nil, nil}
	r4 := &TreeNode{4, nil, nil}
	r5 := &TreeNode{5, nil, nil}
	r6 := &TreeNode{6, nil, nil}
	r7 := &TreeNode{7, nil, nil}

	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r3.Right = r5
	r4.Left = r6
	r5.Right = r7

	fmt.Print(isBalanced(r1))
}

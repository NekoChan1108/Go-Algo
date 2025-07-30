// 自己写的对称二叉树
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树对称是镜像对称的
//换句话说同一高度对应位置左半边节点的右半边要和右半边节点的左半边相同
//右半边节点的左半边要和左半边节点的右半边相同
func IsSame(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSame(p.Left, q.Right) && IsSame(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {

	return IsSame(root.Left, root.Right)
}

func main() {
	// r1 := &TreeNode{1, nil, nil}
	// r2 := &TreeNode{2, nil, nil}
	// r3 := &TreeNode{2, nil, nil}
	// r4 := &TreeNode{4, nil, nil}
	// r5 := &TreeNode{5, nil, nil}
	// r6 := &TreeNode{5, nil, nil}
	// r7 := &TreeNode{4, nil, nil}
	// r1.Left = r2
	// r1.Right = r3
	// r2.Left = r4
	// r2.Right = r5
	// r3.Left = r6
	// r3.Right = r7
	r1 := &TreeNode{1, nil, nil}
	r2 := &TreeNode{2, nil, nil}
	r3 := &TreeNode{2, nil, nil}
	r4 := &TreeNode{3, nil, nil}
	r5 := &TreeNode{3, nil, nil}
	// r6 := &TreeNode{2, nil, nil}
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	// r2.Right = r5
	r3.Left = r5
	fmt.Print(isSymmetric(r1))
}

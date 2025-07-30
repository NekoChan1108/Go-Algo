//官方递归

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//当前树根节点为空节点则表示遍历到叶节点之后了还是找不到p q直接返回空
	if root == nil {
		return nil
	}
	//当前树根节点为二者其一则直接返回
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//遍历左子树
	left := lowestCommonAncestor(root.Left, p, q)
	//遍历右子树
	right := lowestCommonAncestor(root.Right, p, q)
	//如果遍历的结果都不空则表示p q分别在左右子树此时就直接返回当前树根
	if left != nil && right != nil {
		return root
	}
	//如果只在右子树找到就返回右子树根(p q都不在左子树)
	if left == nil {
		return right
	}
	//否则返回左子树根
	return left
}

func main() {
	r1 := &TreeNode{3, nil, nil}
	r2 := &TreeNode{5, nil, nil}
	r3 := &TreeNode{1, nil, nil}
	r4 := &TreeNode{6, nil, nil}
	r5 := &TreeNode{2, nil, nil}
	r6 := &TreeNode{0, nil, nil}
	r7 := &TreeNode{8, nil, nil}
	r8 := &TreeNode{7, nil, nil}
	r9 := &TreeNode{4, nil, nil}
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r2.Right = r5
	r3.Left = r6
	r3.Right = r7
	r5.Left = r8
	r5.Right = r9
	fmt.Println(lowestCommonAncestor(r1, r7, r8).Val)
}

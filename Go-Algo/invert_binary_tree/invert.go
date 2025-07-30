// 自己写的翻转二叉树
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PreOrder(root.Left)
	fmt.Print(root.Val)
	PreOrder(root.Right)
}

//交换
func Swap(p *TreeNode) {
	if p == nil {
		return
	}
	if p.Left == nil && p.Right == nil {
		return
	}
	//先交换左右子树
	// tmp := p.Left
	// p.Left = p.Right
	// p.Right = tmp
	p.Left, p.Right = p.Right, p.Left
	//再交换左右子树的左右子树
	Swap(p.Left)
	Swap(p.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	Swap(root)
	return root
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
	r2.Right = r5
	r3.Left = r6
	r3.Right = r7
	PreOrder(r1)
	fmt.Println()
	PreOrder(invertTree(r1))
}

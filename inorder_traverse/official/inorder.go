// 官方题解(迭代法)
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
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

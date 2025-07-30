// 官方解 队列迭代
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
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

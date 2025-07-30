// 自己写的层序遍历
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	//用于返回
	order := make([][]int, 0)
	//用于存节点
	nodeSlice := make([]*TreeNode, 0)
	if root == nil {
		return order
	}
	nodeSlice = append(nodeSlice, root)
	for len(nodeSlice) > 0 {
		n := len(nodeSlice)
		//用于记录每层的顺序
		valSlice := make([]int, 0)
		for i := 0; i < n; i++ {
			valSlice = append(valSlice, nodeSlice[i].Val)
			if nodeSlice[i].Left != nil {
				nodeSlice = append(nodeSlice, nodeSlice[i].Left)
			}
			if nodeSlice[i].Right != nil {
				nodeSlice = append(nodeSlice, nodeSlice[i].Right)
			}
		}
		nodeSlice = nodeSlice[n:]
		order = append(order, valSlice)
	}
	return order
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
	fmt.Print(levelOrder(r1))
}

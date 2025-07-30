// 官方解(非遍历实现)
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
	//stack记录遍历过的节点
	var stack = make([]*TreeNode, 0)
	node := root
	//存遍历顺序
	var nodeSlice = make([]int, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			nodeSlice = append(nodeSlice, node.Val)
			//先序遍历从根节点一直左走
			node = node.Left
		}
		//走到叶节点的左节点为空时取出stack里存的最后一个也就是左节点为空的那个叶节点
		//让node开始从这个叶节点的右节点开始遍历
		node = stack[len(stack)-1].Right
		//为了保证每个节点访问一次左右都走了直接出栈该节点
		stack = stack[:len(stack)-1]
	}
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

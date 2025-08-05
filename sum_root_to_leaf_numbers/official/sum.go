//官方深度优先遍历

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func main() {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node5 := &TreeNode{5, nil, nil}
	node6 := &TreeNode{6, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node8 := &TreeNode{8, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node4.Left = node8
	node4.Right = node9
	fmt.Println(sumNumbers(node1))
}

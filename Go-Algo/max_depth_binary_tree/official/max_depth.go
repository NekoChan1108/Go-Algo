// 官方题解队列实现
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		//记录当前层的节点个数
		sz := len(queue)
		//把队列里存的当前层节点的子节点存入队列
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		//当前层所有的节点走完就把高度加1
		ans++
	}
	return ans
}
func main() {
	r1 := &TreeNode{1, nil, nil}
	r2 := &TreeNode{2, nil, nil}
	r3 := &TreeNode{3, nil, nil}
	r4 := &TreeNode{4, nil, nil}
	r5 := &TreeNode{5, nil, nil}
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r4.Left = r5
	fmt.Print(maxDepth(r1))
}

// 官方解和自己的一样 这里列出评论区给出的双向列表方法真正实现无需reverse
package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	//返回结果
	res := [][]int{}
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return append(res, []int{root.Val})
	}
	//控制进队出队
	queue := list.New()
	queue.PushBack(root)
	//控制遍历方向(true是从左到右)
	flag := true
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := []int{}
		for i := 0; i < length; i++ {
			if flag {
				//从左往右取
				node := queue.Front().Value.(*TreeNode)
				queue.Remove(queue.Front())
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
				tmp = append(tmp, node.Val)
			} else {
				//从右往左取
				node := queue.Back().Value.(*TreeNode)
				queue.Remove(queue.Back())
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
				tmp = append(tmp, node.Val)
			}
		}
		res = append(res, tmp)
		flag = !flag
	}
	return res
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
	fmt.Print(zigzagLevelOrder(r1))
}

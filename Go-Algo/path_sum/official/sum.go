// 官方双队列解
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		//遍历到叶节点判断是否为预期和
		if now.Left == nil && now.Right == nil {
			//是就返回true
			if temp == sum {
				return true
			}
			//否则继续
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
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
	fmt.Println(hasPathSum(r1, 7))
}

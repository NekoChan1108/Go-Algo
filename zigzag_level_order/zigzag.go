// 自己写的二叉树锯式遍历
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Reverse(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	return nums
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
	queue := []*TreeNode{}
	queue = append(queue, root)
	length := len(queue)

	for len(queue) > 0 {
		//中间结果存每一层的遍历顺序
		tmp := []int{}
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		//循环结束出队前一层的节点
		queue = queue[length:]
		//前一层的遍历结果存到res
		res = append(res, tmp)
		//重新计算长度
		length = len(queue)
	}
	for idx, nums := range res {
		if idx%2 == 1 {
			Reverse(nums)
		}
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

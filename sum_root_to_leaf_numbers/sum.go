// 自己写的全叶节点路径数字之和
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sequenceSlice [][]int

func sumSlice(sequence [][]int) int {
	if len(sequence) == 0 {
		return 0
	}
	sum := 0
	fmt.Println(sequence)
	for _, s := range sequence {
		res := 0
		for i := 0; i < len(s); i++ {
			pow10 := math.Pow10(len(s) - i - 1)
			res += s[i] * int(pow10)
		}
		fmt.Println(res)
		sum += res
	}
	return sum
}

func getAllSequence(root *TreeNode, path []int) {
	//统计每次到叶节点的序列
	if root == nil {
		return
	}
	//path = append(path, root.Val)
	//避免重复
	newPath := append([]int{}, path...)
	newPath = append(newPath, root.Val)
	if root.Left == nil && root.Right == nil {
		sequenceSlice = append(sequenceSlice, newPath)
		return
	}
	if root.Left != nil {
		getAllSequence(root.Left, newPath)
	}
	if root.Right != nil {
		getAllSequence(root.Right, newPath)
	}
}

func sumNumbers(root *TreeNode) int {
	path := make([]int, 0)
	getAllSequence(root, path)
	res := sumSlice(sequenceSlice)
	//清空防止下一调用统计前一次的序列
	sequenceSlice = [][]int{}
	return res
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

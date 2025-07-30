//自己写的二叉树最近公共祖先

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	if root != nil {
		res = append(res, root)
	}
	if root.Left != nil {
		res = append(res, PreOrder(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, PreOrder(root.Right)...)
	}
	return res
}

func Index(list []*TreeNode, node *TreeNode) int {
	for idx, v := range list {
		if node == v {
			return idx
		}
	}
	return -1
}

func BuildNodeMap(node, parent *TreeNode) map[*TreeNode]*TreeNode {
	nodeMap := map[*TreeNode]*TreeNode{}
	if node == nil {
		return nodeMap
	}
	var DFSMap func(node, parent *TreeNode)
	DFSMap = func(node, parent *TreeNode) {
		if parent == nil {
			//根节点
			nodeMap[node] = node
		} else {
			nodeMap[node] = parent
		}
		if node.Left != nil {
			DFSMap(node.Left, node)
		}
		if node.Right != nil {
			DFSMap(node.Right, node)
		}
	}
	DFSMap(node, nil)
	return nodeMap
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || q == nil || p == nil {
		return nil
	}
	if p == q {
		return p
	}
	//获取最近祖先map
	nodeMap := BuildNodeMap(root, nil)
	//获取前序遍历
	order := PreOrder(root)
	//一致就随便返回一个
	if nodeMap[p] == nodeMap[q] {
		return nodeMap[p]
	} else {
		idx1 := Index(order, p)
		idx2 := Index(order, q)
		if idx1 >= idx2 {
			var idx int
			//先从后往前找直到idx2看看父节点的祖先是否包含q
			for idx = idx1; idx >= idx2; {
				if nodeMap[order[idx]] == q {
					return q
				}
				idx = Index(order, nodeMap[order[idx]])
			}
			//不包含先看最后的索引idx对应的节点和q的祖先是否相同
			//不相同q继续向前寻根
			for i := idx2; i >= idx; {
				if order[idx] == nodeMap[order[i]] {
					return order[idx]
				}
				i = Index(order, nodeMap[order[i]])
			}
			return root
		} else {
			//先从后往前找直到idx1看看父节点的祖先是否包含p
			var idx int
			for idx = idx2; idx >= idx1; {
				if nodeMap[order[idx]] == p {
					return p
				}
				idx = Index(order, nodeMap[order[idx]])
			}
			//不包含先看最后的索引idx对应的节点和q的祖先是否相同
			//不相同p继续向前寻根
			for i := idx1; i >= idx; {
				if order[idx] == nodeMap[order[i]] {
					return order[idx]
				}
				i = Index(order, nodeMap[order[i]])
			}
			return root
		}
	}
}

func main() {
	r1 := &TreeNode{3, nil, nil}
	r2 := &TreeNode{5, nil, nil}
	r3 := &TreeNode{1, nil, nil}
	r4 := &TreeNode{6, nil, nil}
	r5 := &TreeNode{2, nil, nil}
	r6 := &TreeNode{0, nil, nil}
	r7 := &TreeNode{8, nil, nil}
	r8 := &TreeNode{7, nil, nil}
	r9 := &TreeNode{4, nil, nil}
	r1.Left = r2
	r1.Right = r3
	r2.Left = r4
	r2.Right = r5
	r3.Left = r6
	r3.Right = r7
	r5.Left = r8
	r5.Right = r9
	fmt.Println(BuildNodeMap(r1, nil))
	fmt.Println(lowestCommonAncestor(r1, r7, r8).Val)
}

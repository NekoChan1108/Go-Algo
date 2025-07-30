//官方递归

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildParentMap(node *TreeNode) map[*TreeNode]*TreeNode {
	parentMap := map[*TreeNode]*TreeNode{}
	if node == nil {
		return parentMap
	}
	var DFS func(node, parent *TreeNode)
	DFS = func(node, parent *TreeNode) {
		if parent == nil {
			//root(最顶上的)节点的父节点就是其本身
			parentMap[node] = node
		} else {
			parentMap[node] = parent
		}
		if node.Left != nil {
			DFS(node.Left, node)
		}
		if node.Right != nil {
			DFS(node.Right, node)
		}
	}
	DFS(node, nil)
	return parentMap
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//当前树根节点为空节点则表示遍历到叶节点之后了还是找不到p q直接返回空
	if root == nil {
		return nil
	}
	//当前树根节点为二者其一则直接返回
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//获取父节点map
	parentMap := BuildParentMap(root)
	//构造访问map
	visitMap := map[*TreeNode]bool{}
	//!=root防止死循环
	for p != root {
		visitMap[p] = true
		p = parentMap[p]
	}
	for q != root {
		//q遍历到p走过的父节点就返回这个最近的公共祖先
		if visitMap[q] {
			return q
		}
		q = parentMap[q]
	}
	//二者其一走到root说明最近的只能是root
	if q == root || p == root {
		return root
	}
	return nil
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
	// fmt.Println(BuildParentMap(r1))
	fmt.Println(lowestCommonAncestor(r1, r7, r8).Val)
}

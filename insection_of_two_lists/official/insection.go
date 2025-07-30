// 官方题解(哈希map)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//记录节点
	nodeMap := make(map[*ListNode]bool)
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		nodeMap[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		//遍历到A读取的直接返回
		if nodeMap[tmp] {
			return tmp
		}
	}
	return nil
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l4 := &ListNode{4, nil}
	l1.Next = l2
	l2.Next = l4
	l3.Next = l4
	fmt.Println(*getIntersectionNode(l1, l3))
}

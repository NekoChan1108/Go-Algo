// 自己写的(倒数第k个节点)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	if head == nil {
		return nil
	}
	var i int = 0
	nodeMap := make(map[int]*ListNode)
	for head != nil {
		i++
		nodeMap[i] = head
		head = head.Next
	}
	if cnt > i {
		return nil
	}
	return nodeMap[i-cnt+1]
}

func main() {
	ll := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Print(trainingPlan(ll, 4))
}

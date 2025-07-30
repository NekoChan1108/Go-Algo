// 官方没有go(哈希map解法)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	//记录遍历过的节点
	nodeMap := make(map[*ListNode]bool)
	if head == nil || head.Next == nil {
		return false
	}
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		if head.Next != nil {
			nodeMap[head] = true
			head = head.Next
		} else {
			return false
		}
	}
	return false
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l2
	fmt.Println(hasCycle(l1))
}

// 官方哈希表解
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]bool{}
	//一直遍历直到遇到第一个已经遍历过存在map里的节点
	for head != nil {
		if val, ok := seen[head]; ok && val {
			return head
		}
		seen[head] = true
		head = head.Next
	}
	return nil
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l1
	fmt.Print(*detectCycle(l1))
}

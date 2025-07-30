// 自己写的删除链表
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//存索引到map 用length记录链表长度 res用于遍历
	nodeMap, length, res := map[int]*ListNode{}, 0, head
	for res != nil {
		length++
		nodeMap[length] = res
		res = res.Next
	}
	fmt.Printf("length: %v\n", length)
	if n > length {
		return nil
	}
	//拿到倒数n+1 n
	var node1, node2 *ListNode
	if n == length {
		return head.Next
	}
	node1 = nodeMap[length-n+1]
	node2 = nodeMap[length-n]
	if n == 1 {
		node2.Next = nil
		return head
	}
	node2.Next = node1.Next
	node1.Next = nil
	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l4 := &ListNode{4, nil}
	l5 := &ListNode{5, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	node := removeNthFromEnd(l1, 2)
	for node != nil {
		fmt.Print(node.Val )
		node = node.Next
	}
}

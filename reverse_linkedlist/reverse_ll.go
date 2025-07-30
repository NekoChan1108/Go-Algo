// 反转链表(自己做的)
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//创建一个空节点用于连接反转后的链表
	p := new(ListNode)
	//让q指向head记录head且永远指向反转后的头节点
	q := new(ListNode)
	if head != nil && head.Next == nil {
		return head
	}
	if head == nil {
		return nil
	}
	for head != nil {
		q = head
		head = head.Next
		q.Next = p.Next
		p.Next = q
	}

	return q
}

func main() {
	res := reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{}}}})
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}

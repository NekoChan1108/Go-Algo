// 自己写的删除链表重复元素II
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->2->3->4->4->5 ===> 1->3->5
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	//无论头节点被删多少次dummy永远指向那个新的头节点
	dummy := &ListNode{0, head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			//跳过值为x的节点
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{2, nil}
	// l4 := &ListNode{3, nil}
	// l5 := &ListNode{4, nil}
	// l6 := &ListNode{4, nil}
	// l7 := &ListNode{5, nil}
	l1.Next = l2
	l2.Next = l3
	// l3.Next = l4
	// l4.Next = l5
	// l5.Next = l6
	// l6.Next = l7
	node := deleteDuplicates(l1)
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}

}

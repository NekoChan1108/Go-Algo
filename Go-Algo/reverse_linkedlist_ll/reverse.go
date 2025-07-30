// 自己写的反转链表Ⅱ
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//根据位置查找结点
func Index(head *ListNode, idx int) *ListNode {
	if head == nil {
		return nil
	}
	cnt := 1
	for head != nil {
		if cnt == idx {
			return head
		}
		head = head.Next
		cnt++
	}
	return nil
}

func reverse(head, target *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//创建p、q分别用于连接反转后的节点以及指向反转后的头节点
	p, q := new(ListNode), new(ListNode)
	for head != target {
		q = head
		head = head.Next
		q.Next = p.Next
		p.Next = q
	}
	//最终是一个空节点p连接上反转后的链表
	return q
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left > right || head == nil || head.Next == nil {
		return head
	}

	//定位四个left-1 left right right+1
	p, q, r, s := Index(head, left-1), Index(head, left), Index(head, right), Index(head, right+1)
	// fmt.Printf("p q r s: %v %v %v %v\n", p.Val, q.Val, r.Val, s.Val)
	//记录需要反转部分反转之后的头尾节点
	rTail, rHead := q, r
	//反转
	reverse(q, s)
	//重连接
	if p != nil {
		p.Next = rHead
	} else {
		head.Next = s
		head = rHead
	}
	if s != nil {
		rTail.Next = s

	}

	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l4 := &ListNode{4, nil}
	l5 := &ListNode{5, nil}
	l6 := &ListNode{6, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	for head := reverseBetween(l1, 1, 6); head != nil; head = head.Next {
		fmt.Print(head.Val)
	}
}

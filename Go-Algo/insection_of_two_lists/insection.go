// 自己写的(找出两个列表相交的交点)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB
	l1 := 0
	l2 := 0
	//计算长度
	for p != nil {
		l1++
		p = p.Next
	}
	for q != nil {
		l2++
		q = q.Next
	}
	// fmt.Println(l1, l2)
	p = headA
	q = headB
	//移动到对齐的起点
	if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			p = p.Next
		}
		if p == q {
			return p
		}
	} else {
		for i := 0; i < l2-l1; i++ {
			q = q.Next
		}
		if p == q {
			return p
		}
	}
	//同时遍历
	for p != nil && q != nil {
		p = p.Next
		q = q.Next
		if p == q {
			return p
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

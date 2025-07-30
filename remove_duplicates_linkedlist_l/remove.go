// 自己写的移除有序链表的重复元素
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	//如果链表为空或者链表只有一个元素则就是无重复
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	//用于返回
	r := head
	for q != nil {
		if p.Val == q.Val {
			tmp := q.Next
			p.Next = tmp
			q.Next = nil
			q = tmp
		} else {
			p = q
			q = q.Next
		}
	}
	return r
}

func main() {
	ll := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	p := ll
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println()
	nll := deleteDuplicates(ll)
	p = nll
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
}

// 判断链表是否有环(自己写的)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	//快慢指针
	p := head
	q := head
	for p != nil && q != nil {
		if p.Next != nil {
			p = p.Next
		} else {
			return false
		}
		if q.Next != nil && q.Next.Next != nil {
			q = q.Next.Next
		} else {
			return false
		}
		if p == nil || q == nil {
			break
		}
		if p == q {
			return true
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
	// l3.Next=l2
	fmt.Println(hasCycle(l1))
}

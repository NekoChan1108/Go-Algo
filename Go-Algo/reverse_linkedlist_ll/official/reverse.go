// 官方一解和自己一样二解只需遍历一次
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
    // 设置 dummyNode 是这一类问题的一般做法
    dummyNode := &ListNode{Val: -1}
    dummyNode.Next = head
    pre := dummyNode
    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }
    cur := pre.Next
    for i := 0; i < right-left; i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return dummyNode.Next
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

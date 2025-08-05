// 官方两种归并排序
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

/*
*
自顶向下归并排序
*/
func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

/**
自底向上归并排序
*/
//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	length := 0
//	for node := head; node != nil; node = node.Next {
//		length++
//	}
//
//	dummyHead := &ListNode{Next: head}
//	for subLength := 1; subLength < length; subLength <<= 1 {
//		prev, cur := dummyHead, dummyHead.Next
//		for cur != nil {
//			head1 := cur
//			for i := 1; i < subLength && cur.Next != nil; i++ {
//				cur = cur.Next
//			}
//
//			head2 := cur.Next
//			cur.Next = nil
//			cur = head2
//			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
//				cur = cur.Next
//			}
//
//			var next *ListNode
//			if cur != nil {
//				next = cur.Next
//				cur.Next = nil
//			}
//
//			prev.Next = merge(head1, head2)
//
//			for prev.Next != nil {
//				prev = prev.Next
//			}
//			cur = next
//		}
//	}
//	return dummyHead.Next
//}

func main() {
	l1 := &ListNode{-3, nil}
	l2 := &ListNode{-5, nil}
	l3 := &ListNode{-7, nil}
	l4 := &ListNode{7, nil}
	l5 := &ListNode{5, nil}
	l6 := &ListNode{3, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6

	sortL := sortList(l1)

	for sortL != nil {
		fmt.Println(sortL.Val)
		sortL = sortL.Next
	}
}

// 官方解Ⅰ和自己写的差不多解Ⅱ实现了O(1)复杂度
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//利用快慢指针找到中间节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//快指针走完了为nil此时慢指针指向的就是len(list)/2处
	return slow
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

//交替合并列表
func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next
		l1.Next = l2
		l1 = l1Tmp
		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
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
	reorderList(l1)
	for head := l1; head != nil; head = head.Next {
		fmt.Print(head.Val)
	}
}

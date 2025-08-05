// 自己写的排序链表
package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodeValSlice := make([]*ListNode, 0)
	for head != nil {
		nodeValSlice = append(nodeValSlice, head)
		head = head.Next
	}
	sort.Slice(nodeValSlice, func(i, j int) bool {
		return nodeValSlice[i].Val <= nodeValSlice[j].Val
	})
	for idx := range nodeValSlice {
		if idx == len(nodeValSlice)-1 {
			nodeValSlice[idx].Next = nil
		} else {
			nodeValSlice[idx].Next = nodeValSlice[idx+1]
		}
	}
	return nodeValSlice[0]
}

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

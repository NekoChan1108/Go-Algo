// 合并两个有序列表(自己写的 官方没有go🤣)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//用一个空节点P来移动定位最小的节点
	p := new(ListNode)
	//q记录初始的空节点返回q.next就是结果
	q := new(ListNode)
	q = p
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			p = list1
			list1 = list1.Next
			if list1 != nil && list1.Val >= list2.Val {
				p.Next = list2
				p = list2
				list2 = list2.Next
			}
		} else {
			p.Next = list2
			p = list2
			list2 = list2.Next
			if list2 != nil && list2.Val >= list1.Val {
				p.Next = list1
				p = list1
				list1 = list1.Next
			}
		}

	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return q.Next
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	list2 := &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}
	res := mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}

// 自己写的环形链表2
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	//快慢指针
	slow, fast := head, head
	//奇数个定位到最后一个节点 偶数个定位到最后一个节点的下一个
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	// 找到环入口
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l1
	fmt.Print(*detectCycle(l1))
}

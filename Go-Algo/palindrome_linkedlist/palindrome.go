// 自己写的(回文链表)
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var nodeMap = make(map[int]*ListNode, 0)
	var idx, len int
	for head != nil {
		nodeMap[idx] = head
		head = head.Next
		idx++
	}
	len = idx
	for i := 0; i < len/2; i++ {
		if nodeMap[i] != nil && nodeMap[len-i-1] != nil && nodeMap[i].Val != nodeMap[len-i-1].Val {
			return false
		}
	}
	return true
}

func main() {
	ll := &ListNode{1, &ListNode{2, nil}}
	fmt.Print(isPalindrome(ll))
}

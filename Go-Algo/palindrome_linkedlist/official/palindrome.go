// 官方解(存数组比较前后两半)
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
	var nodeSlice=make([]int, 0)
	for head != nil {
		nodeSlice = append(nodeSlice, head.Val)
		head=head.Next
	}
	
	for i := 0; i < len(nodeSlice)/2; i++ {
		if nodeSlice[i]!=nodeSlice[len(nodeSlice)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	ll := &ListNode{1, &ListNode{2, nil}}
	fmt.Print(isPalindrome(ll))
}

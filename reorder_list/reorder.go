// 自己写的重排列表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//返回长度以及各个下表与节点的对应关系
func IndexMap(head *ListNode) (map[int]*ListNode, int) {
	if head == nil {
		return nil, 0
	}
	length := 0
	nodeMap := map[int]*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodeMap[length] = cur
		length++
	}
	return nodeMap, length
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	nodeMap, length := IndexMap(head)
	fmt.Printf("nodeMap: %v\n", nodeMap)
	fmt.Printf("length: %v\n", length)
	//只需遍历一半就可以
	var i int
	for i = 0; i < length/2; i++ {
		nodeMap[length-i-1].Next = nodeMap[i].Next
		nodeMap[i].Next = nodeMap[length-i-1]
	}
	fmt.Printf("当前节点：%v\n", nodeMap[i])
	nodeMap[i].Next = nil
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l4 := &ListNode{4, nil}
	l5 := &ListNode{5, nil}
	// l6 := &ListNode{6, nil}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	// l5.Next = l6
	reorderList(l1)
	for head := l1; head != nil; head = head.Next {
		fmt.Print(head.Val)
	}
}

// 自己写的删除链表重复元素II
package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->2->3->4->4->5 ===> 1->3->5
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//记录节点第一次出现的索引以及节点出现的次数
	idxMap, cntMap, nodeMap, length, res, dummy := map[int]int{}, map[int]int{}, map[int]*ListNode{}, 0, head, &ListNode{0, nil}

	for res != nil {
		length++
		//如果没有索引就加入有索引代表重复出现无需更新
		if _, ok := idxMap[res.Val]; !ok {
			idxMap[res.Val] = length
			//加入节点记录第一次出现的节点在哪
			nodeMap[length] = res
		}
		cntMap[res.Val]++
		res = res.Next
	}
	fmt.Println(idxMap)
	fmt.Println(cntMap)
	fmt.Println(nodeMap)
	//遍历map拿到出现次数等于1的节点
	for k, v := range cntMap {
		if v > 1 {
			//删掉次数大于1的
			delete(cntMap, k)
			//删掉nodeMap里次数大于1的节点
			delete(nodeMap, idxMap[k])
		}
	}
	fmt.Println(idxMap)
	fmt.Println(cntMap)
	fmt.Println(nodeMap)
	//拼接链表
	//按key排序nodeMap
	keys := make([]int, 0)
	for k := range nodeMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	p := dummy
	for _, key := range keys {
		dummy.Next = nodeMap[key]
		dummy = nodeMap[key]
	}
	//截断
	dummy.Next = nil
	return p.Next
}

func main() {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{2, nil}
	// l4 := &ListNode{3, nil}
	// l5 := &ListNode{4, nil}
	// l6 := &ListNode{4, nil}
	// l7 := &ListNode{5, nil}
	l1.Next = l2
	l2.Next = l3
	// l3.Next = l4
	// l4.Next = l5
	// l5.Next = l6
	// l6.Next = l7
	node := deleteDuplicates(l1)
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}

}

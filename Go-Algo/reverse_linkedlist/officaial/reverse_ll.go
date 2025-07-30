// 官方题解(第一种和自己的一样 这里采用第二种递归)
package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//递归的结束条件
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
	//反转的本质是节点的下一个要变成节点的上一个
    head.Next.Next = head
    head.Next = nil
    return newHead
}

func main(){
	res:=reverseList(&ListNode{1,&ListNode{2,&ListNode{3,&ListNode{}}}})
	for res!=nil{
		fmt.Print(res.Val)
		res=res.Next
	}
}
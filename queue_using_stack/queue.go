// 用两个栈实现队列(自己写的)
package main

import "fmt"

type MyQueue struct {
	S1 []int
	S2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		//s1进
		S1: make([]int, 0),
		//s2接收s1出
		S2: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.S1 = append(q.S1, x)
}

func (q *MyQueue) PushS2() {
	for len(q.S1) > 0 {
		q.S2 = append(q.S2, q.S1[len(q.S1)-1])
		q.S1 = q.S1[:len(q.S1)-1]
	}
}

func (q *MyQueue) Pop() int {
	//取不出从s1拿
	if len(q.S2) == 0 {
		q.PushS2()
	}
	val := q.S2[len(q.S2)-1]
	q.S2 = q.S2[:len(q.S2)-1]
	return val
}

func (q *MyQueue) Peek() int {
	//取不出就从s1拿
	if len(q.S2) == 0 {
		q.PushS2()
	}
	return q.S2[len(q.S2)-1]
}

func (q *MyQueue) Empty() bool {
	//两个栈全为0就空
	return len(q.S1) == 0 && len(q.S2) == 0
}

func main() {
	q := &MyQueue{}
	for i := 1; i <= 5; i++ {
		q.Push(i)
	}
	fmt.Println(q.S1, q.S2)
	for i := 1; i <= 5; i++ {
		fmt.Printf("出队：%v\n", q.Pop())
	}
	fmt.Println(q.S1, q.S2)
	fmt.Println(q.Empty())
}

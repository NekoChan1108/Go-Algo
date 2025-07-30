// 自己写的(最小栈)
package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	//同步存最小值
	min []int
}

func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), min: []int{math.MaxInt}}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	s.min = append(s.min, min(val, s.min[len(s.min)-1]))
}

func (s *MinStack) Pop() {
	s.min = s.min[:len(s.min)-1]
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

// 常数时间获取
func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}

func main() {
	s := Constructor()
	s.Push(2)
	s.Push(0)
	s.Push(3)
	s.Push(0)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.GetMin())
}

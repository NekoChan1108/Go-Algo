// 自己写的将数组的0元素全部移动到末尾同时保持非0元素的顺序
package main

import "fmt"

func moveZeroes(nums []int) {
	// fmt.Println(nums)
	//存0
	numsTmp1 := make([]int, 0)
	//存非0
	numsTmp2 := make([]int, 0)

	for _, v := range nums {
		if v != 0 {
			numsTmp2 = append(numsTmp2, v)
		} else {
			numsTmp1 = append(numsTmp1, v)
		}
	}
	copy(nums, append(numsTmp2, numsTmp1...))
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{1, 0, 2, 3, 0})
}

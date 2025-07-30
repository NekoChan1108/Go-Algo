// 自己写的多数元素
package main

import "fmt"

func majorityElement(nums []int) int {
	n := len(nums) / 2
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	for key, value := range cnt {
		if value > n {
			return key
		}
	}
	return 0
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

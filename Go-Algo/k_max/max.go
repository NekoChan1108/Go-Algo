// 自己写的找出数组第K个最大数
package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k < 1 {
		return 0
	}
	if len(nums) == 1 && k == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}
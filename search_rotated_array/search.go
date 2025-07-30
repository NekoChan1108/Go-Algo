// 自己写的在随机旋转后的数组查找对应目标的下标
package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 || (len(nums) == 1 && target != nums[0]) ||
		(target > nums[len(nums)-1] && target < nums[0]) {
		return -1
	}
	if target == nums[0] || (target == nums[0] && len(nums) == 1) {
		return 0
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	for idx, v := range nums {
		if v == target {
			return idx
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Print(search(nums, 0))
}

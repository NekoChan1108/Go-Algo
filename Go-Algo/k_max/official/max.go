// 官方解快排
package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		//向右移动找到第一个比pivot大的
		for i++; nums[i] < partition; i++ {
		}
		//向左移动找到第一个比pivot小的
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		//值<=pivot的那左半边
		return quickselect(nums, l, j, k)
	} else {
		//值>pivot的那半边
		return quickselect(nums, j+1, r, k)
	}
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}

// 官方解二分法
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
	var left, right = 0, len(nums) - 1
	for left <= right {
		//二分一定是有一半有序的，先去有序的那部分里查，查不到再去无序的里面继续二分
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		//判断哪一边是有序的
		if nums[0] < nums[mid] {
			//左半边有序
			//target在左半边范围
			if target < nums[mid] && target >= nums[0] {
				right = mid - 1
			} else {
				//target不在左半边范围
				left = mid + 1
			}
		} else {
			//右半边有序
			//target在右半边范围
			if target < nums[len(nums)-1] && target >= nums[mid] {
				left = mid + 1
			} else {
				//target不在右半边范围
				right = mid - 1
			}

		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Print(search(nums, 0))
}

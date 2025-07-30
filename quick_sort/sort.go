// 自己写的快排
package main

import (
	"fmt"
)

func IsASCSorted(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}
		if nums[i] > nums[i+1] && i+1 < len(nums) {
			return false
		}
	}
	return true
}

func IsDESCSorted(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}
		if nums[i] < nums[i+1] && i+1 < len(nums) {
			return false
		}
	}
	return true
}

func Reverse(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	return nums
}

// 分区返回下标
func Partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		//从右往左找到第一个比pivot小的
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		//从左往右找第一个比pivot大的
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[right] = pivot
	return left
}

func QuickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	//分治
	idx := Partition(nums, left, right)
	QuickSort(nums, left, idx-1)
	QuickSort(nums, idx+1, right)
}

// [1,21,100,89,3,50]
func sortArray(nums []int) []int {
	if IsASCSorted(nums) {
		return nums
	}
	if IsDESCSorted(nums) {
		return Reverse(nums)
	}
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	nums := []int{21, 100, 3, 50, 1}
	fmt.Println(sortArray(nums))
	fmt.Println(Reverse(nums))
}

// 官方解随机抽取
package main

import (
	"fmt"
	"math/rand"
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

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	idx := RandomPartition(nums, left, right)
	QuickSort(nums, left, idx-1)
	QuickSort(nums, idx+1, right)
}

func RandomPartition(nums []int, left, right int) int {
	//随机选取主元
	idx := rand.Intn(right-left+1) + left
	// fmt.Printf("随机下标为: %v\n", idx)
	//和最左边的交换位置
	// fmt.Printf("交换前: %v\n", nums[left])
	nums[left], nums[idx] = nums[idx], nums[left]
	// fmt.Printf("交换后: %v\n", nums[left])
	//取得主元
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

func main() {
	nums := []int{21, 100, 3, 50, 1}
	fmt.Println(sortArray(nums))
	fmt.Println(Reverse(nums))
}

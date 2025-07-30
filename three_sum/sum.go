// 自己写的三数之和
package main

import (
	"fmt"
	"sort"
)

func Sum(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum == 0
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	//排序
	sort.Ints(nums)
	if len(nums) < 3 || (len(nums) == 3 && !Sum(nums)) {
		return [][]int{}
	}
	if len(nums) == 3 && Sum(nums) {
		numSlice := make([][]int, 0)
		return append(numSlice, nums)
	}
	for i := 0; i < len(nums)-2; i++ {
		//双指针
		if i > 0 && nums[i] == nums[i-1] {
			//去重第一个
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重第二个、第三个数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
				//不够就加大的往右走
			} else if sum < 0 {
				left++
				//大了就加小的往左走
			} else {
				right--
			}
		}

	}
	return res
}

func main() {
	// nums := []int{-1, 1, -2, 2}
	// fmt.Print(TwoSum(nums, 0))
	nums := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}
	fmt.Println(threeSum(nums))
}

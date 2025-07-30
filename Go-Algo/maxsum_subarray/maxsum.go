// 自己写的最大子数组和(超时)
package main

import (
	"fmt"
)

// func GetMaxFromSlice(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	max := math.MinInt
// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }

// func GetMaxFromSlice2(sumSlice [][]int) int {
// 	maxSlice := make([]int, 0)
// 	if len(sumSlice) == 0 {
// 		return 0
// 	}
// 	if len(sumSlice) == 1 {
// 		return GetMaxFromSlice(sumSlice[0])
// 	}
// 	for _, v := range sumSlice {
// 		maxSlice = append(maxSlice, GetMaxFromSlice(v))
// 	}
// 	return GetMaxFromSlice(maxSlice)
// }

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// sumSlice := make([][]int, 0)
	n := len(nums)
	maxSum := nums[0]
	for i := 0; i < n; i++ {
		// tmpSum := nums[i]
		tmpSum := 0
		// tmpSlice := []int{}
		// tmpSlice = append(tmpSlice, nums[i])
		for j := i; j < n; j++ {
			tmpSum += nums[j]
			// tmpSlice = append(tmpSlice, tmpSum)
			if tmpSum > maxSum {
				maxSum = tmpSum
			}
		}
		// sumSlice = append(sumSlice, tmpSlice)
	}
	// //得到所有子序列的值
	// fmt.Println(sumSlice)
	// return GetMaxFromSlice2(sumSlice)
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

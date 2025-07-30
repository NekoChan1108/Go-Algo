// 两数之和(自己做的)
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	idx := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				idx = append(idx, i, j)
				return idx
			}
		}
	}
	return idx
}

func main() {
	idx := twoSum([]int{1, 2, 3, 4}, 6)
	fmt.Println(idx)
}

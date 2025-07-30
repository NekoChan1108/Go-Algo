// 官方解法双指针
package main

import "fmt"

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			//不为0就交换左右元素把不为0的换到前面去
			//这一句其实就是交换两者相当于tmp:=nums[left] nums[left]=nums[right] nums[right]=tmp
			nums[left], nums[right] = nums[right], nums[left]
			//左指针有移动继续寻找0
			left++
		}
		//为0右指针继续移动寻找非0
		right++
	}
	fmt.Print(nums)
}

func main() {
	moveZeroes([]int{1, 0, 2, 3, 0})
}

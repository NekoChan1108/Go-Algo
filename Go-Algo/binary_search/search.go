// 自己写的(二分查找)
package main

import "fmt"

func search(nums []int, target int) int {
	var i int = 0
	var j int = len(nums) - 1
	var k int
	for j >= i {
		k = (i + j) / 2
		if nums[k]==target{
			return k
		}
		if target < nums[k] {
			j=k-1
		}else if target > nums[k]{
			i=k+1
		}
	}
	return -1
}

func main(){
	nums := []int{5}
	fmt.Println(search(nums,5))
}
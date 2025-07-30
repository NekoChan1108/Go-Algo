//自己写的全排列

package main

import (
	"fmt"
)

func BackTrack(first int, nums []int) [][]int {
	res := [][]int{}
	//递归终止条件 first决定当前位置上的元素顺序剩下的交给之后处理
	//first==len(nums)说明所有元素都排好了可以保存序列了
	if first == len(nums) {
		//复制当前状态防止后续影响
		res = append(res, append([]int(nil), nums...))
		return res
	}
	//从first而不是从0开始是因为前first-1个顺序已经固定好了再从0开始会造成重复
	for idx := first; idx < len(nums); idx++ {
		//将 nums[idx] 作为当前 first 位置的元素（例如 [1,2,3] ➝ [2,1,3]）
		nums[first], nums[idx] = nums[idx], nums[first]
		//固定 first 位置后，递归生成后续所有排列，并将结果合并到 res 中
		res = append(res, BackTrack(first+1, nums)...)
		//回溯：撤销交换，恢复 nums 状态，以便尝试下一次循环的交换
		nums[first], nums[idx] = nums[idx], nums[first]
	}
	return res
}

func permute(nums []int) [][]int {
	return BackTrack(0, nums)
}

func main() {
	fmt.Print(permute([]int{1, 2, 3, 4}))
}

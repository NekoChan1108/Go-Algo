// 合并两个有序数组(自己写的)
package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	merged := make([]int, m+n)
	i := 0
	j := 0
	k := 0
	//归并排序(切片底层采用拷贝copy方法才能实现值覆盖)
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			merged[k] = nums1[i]
			i++
			k++
		} else {
			merged[k] = nums2[j]
			j++
			k++
		}
	}
	for i < m {
		merged[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		merged[k] = nums2[j]
		k++
		j++
	}
	copy(nums1, merged)
}
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 3, 4}
	merge(nums1, 3, nums2, 2)
	fmt.Println(nums1)
}

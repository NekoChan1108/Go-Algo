//官方解法
package main

import (
	"fmt"
	"sort"
)
func merge(nums1 []int, m int, nums2 []int, n int) {
    copy(nums1[m:], nums2)
    sort.Ints(nums1)
}
func main(){
	nums1:=[]int{1,2,3,0,0,0}
	nums2:=[]int{2,3,4}
	merge(nums1,3,nums2,2)
	fmt.Println(nums1)
}
// 两数之和官方题解
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, val := range nums {
		if j, ok := hashTable[target-val]; ok {
			return []int{j, i}
		}
		hashTable[val] = i
	}
	return nil
}

func main() {
	idx := twoSum([]int{1, 2, 3, 4}, 6)
	fmt.Println(idx)
}

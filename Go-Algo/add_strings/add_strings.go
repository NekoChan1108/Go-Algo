// 自己写的
package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	l3 := max(l1, l2) + 1
	res := make([]int, l3)
	var i = l1 - 1
	var j = l2 - 1
	var k = l3 - 1
	var add int = 0
	for i >= 0 && j >= 0 && k >= 0 {
		//int强转unit8是assiic值减去对应'0'的assiic值就是结果
		val1 := int(num1[i]) - int('0')
		val2 := int(num2[j]) - int('0')
		val := val1 + val2 + add
		if val >= 10 {
			res[k] = val - 10
			add = 1
		} else {
			res[k] = val
			add = 0
		}
		i--
		j--
		k--
	}
	if j < 0 {
		for ; i >= 0; i-- {
			val := int(num1[i]) - int('0') + add
			if val >= 10 {
				res[k] = val - 10
				add = 1
			} else {
				res[k] = val
				add = 0
			}
			k--
		}
	}
	if i < 0 {
		for ; j >= 0; j-- {
			val := int(num2[j]) - int('0') + add
			if val >= 10 {
				res[k] = val - 10
				add = 1
			} else {
				res[k] = val
				add = 0
			}
			k--
		}
	}
	if i < 0 && j < 0 {
		res[k] = add
	}
	if res[k] == 0 {
		res = res[1:]
	}
	var addStr string
	for idx := 0; idx < len(res); idx++ {
		addStr += strconv.Itoa(res[idx])
	}
	return addStr
}

func main() {
	fmt.Print(addStrings("6", "501"))
}

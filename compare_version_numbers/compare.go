// 自己写的比较版本大小
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	//获取按.分割版本号的数组
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	fmt.Println(v1)
	fmt.Println(v2)
	length := 0
	//不足就填充0
	if len(v1) < len(v2) {
		length = len(v1)
		for i := 0; i < len(v2)-length; i++ {
			v1 = append(v1, "0")
		}
	} else {
		length = len(v2)
		for i := 0; i < len(v1)-length; i++ {
			v2 = append(v2, "0")
		}
	}
	length = len(v1)
	fmt.Println(v1)
	fmt.Println(v2)
	//遍历对齐之后的v1 v2去除多余的0
	for i := 0; i < length; i++ {
		for j := 0; j < len(v1[i]); j++ {
			if strings.HasPrefix(v1[i], "0") && len(v1[i]) > 1 {
				v1[i] = v1[i][1:]
				j = 0
			}
		}
		for j := 0; j < len(v2[i]); j++ {
			if strings.HasPrefix(v2[i], "0") && len(v2[i]) > 1 {
				v2[i] = v2[i][1:]
				j = 0
			}
		}
	}
	fmt.Println(v1)
	fmt.Println(v2)
	//再遍历
	for i := 0; i < length; i++ {
		val1, _ := strconv.Atoi(v1[i])
		val2, _ := strconv.Atoi(v2[i])
		if val1 > val2 {
			return 1
		} else if val1 < val2 {
			return -1
		} else {
			continue
		}
	}
	return 0
}
func main() {
	str1 := "1.0"
	str2 := "1.0.0.0"
	//fmt.Println(strings.Split(str1, "."))
	//fmt.Println(strings.Split(str2, "."))
	fmt.Println(compareVersion(str1, str2))
}

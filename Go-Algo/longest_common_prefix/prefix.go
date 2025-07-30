// 自己写的求最长公共前缀
package main

import (
	"fmt"
	"math"
	"strings"
)

func minStr(strs []string) string {
	var minStr = strs[0]
	for _, v := range strs {
		if len(v) < len(minStr) {
			minStr = v
		}
	}
	return minStr
}

func longestCommonPrefix(strs []string) string {

	var minLen, minIdx, tmp int
	var prefix []string
	minLen = math.MaxInt

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	//遍历strs发现有""直接返回
	//先找到最短的那个
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		if len(strs[i]) < minLen {
			minIdx = i
			minLen = len(strs[i])
			tmp = minLen
		}
	}
	// fmt.Println(minIdx, minLen, strs[minIdx])
	//以最短的作为模板遍历strs
	for j := 0; j < len(strs); j++ {
		//只要出现首字母和模板不一样就返回""
		if strs[minIdx][0] != strs[j][0] {
			return ""
		}
		//跳过模板
		if j == minIdx {
			continue
		}
		// fmt.Printf("j:%v\n", j)
		//比较模板和当前遍历到的str的同等长度
		for k := len(strs[j]); k > 0; k-- {
			//tmp一直减小到没有了说明不符合要求直接返回空
			if tmp <= 0 {
				return ""
			}
			// fmt.Printf("tmp:%v\n", tmp)
			val := strings.Compare(strs[j][:tmp], strs[minIdx][:tmp])
			//如果相同就跳出循环继续下一个
			if val == 0 {
				//存入前缀数组
				prefix = append(prefix, strs[minIdx][:tmp])
				break
			}
			//如果不是就缩短继续比较
			tmp--
		}
		//循环结束还原
		tmp = minLen
	}
	// fmt.Println(tmp)
	// fmt.Println(prefix)
	return minStr(prefix)
}

func main() {
	strs := []string{"car", "carc", "carv", "cacrr"}
	fmt.Println(longestCommonPrefix(strs))
}

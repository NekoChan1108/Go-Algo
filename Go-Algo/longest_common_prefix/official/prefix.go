// 官方题解纵向扫描
package main

import (
	"fmt"
)

//纵向扫描时，从前往后遍历所有字符串的每一列，比较相同列上的字符是否相同，
// 如果相同则继续对下一列进行比较，如果不相同则当前列不再属于公共前缀，
// 当前列之前的部分为最长公共前缀。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			//i==len(strs[j])说明此时的strs[j]长度最短直接返回
			//str[j][i]和str[0][i]不相同则返回前i个
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	//如果str[0]都走完了说明str[0]长度最小且为公共前缀
	return strs[0]
}

func main() {
	strs := []string{"car", "carc", "carv", "cacrr"}
	fmt.Print(longestCommonPrefix(strs))
}

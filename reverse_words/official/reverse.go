// 官方无go语言版本 参考评论区双指针实现
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	m := len(s) - 1
	var res strings.Builder

	// 去除尾部空格
	for m >= 0 && s[m] == ' ' {
		m--
	}

	n := m
	for m >= 0 {
		// 找到单词的起始位置
		for m >= 0 && s[m] != ' ' {
			m--
		}
		// 将单词加入结果
		res.WriteString(s[m+1 : n+1])
		res.WriteByte(' ')

		// 跳过空格
		for m >= 0 && s[m] == ' ' {
			m--
		}
		n = m
	}

	// 去掉最后一个多余的空格
	result := res.String()
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}
func main() {
	s := "  hello world kkk  "
	fmt.Println(reverseWords(s))
}

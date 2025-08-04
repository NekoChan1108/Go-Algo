// 自己写的反转字符串中单词的顺序
package main

import (
	"fmt"
	"strings"
)

func reverse(str []string) []string {
	length := len(str)
	if length == 0 || length == 1 {
		return str
	}
	for i := 0; i < length/2; i++ {
		str[i], str[length-i-1] = str[length-i-1], str[i]
	}
	return str
}
func reverseWords(s string) string {
	split := strings.Split(s, " ")
	//fmt.Println(split)
	strWithoutSpace := make([]string, 0)
	resStr := ""
	for _, str := range split {
		if str != "" {
			strWithoutSpace = append(strWithoutSpace, str)
		}
	}
	reverseStrWithoutSpace := reverse(strWithoutSpace)
	for i, str := range reverseStrWithoutSpace {
		resStr += str
		if i != len(reverseStrWithoutSpace)-1 {
			resStr += " "
		}
	}
	return resStr
}

func main() {
	s := "  hello world  "
	fmt.Println(reverseWords(s))
}

// 自己写的无重复最长子串长度(超时)
package main

import "fmt"

func Judege(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	n := 1
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if Judege(s[i:j+1]) && len(s[i:j+1]) > n {
				n = len(s[i : j+1])
			}
		}
	}
	return n
}

func main() {
	fmt.Print(lengthOfLongestSubstring("abcabcbb"))
}

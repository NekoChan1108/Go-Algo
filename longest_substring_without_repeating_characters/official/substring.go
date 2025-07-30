// 官方解滑动窗口
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	charIndex := make(map[byte]int) // 存储字符最后出现的位置
	start := 0                      // 窗口起始位置
	maxLength := 0                  // 最长子串长度

	for end := 0; end < len(s); end++ {
		if idx, found := charIndex[s[end]]; found {
			// 如果字符已存在，移动窗口起始位置到重复字符的下一个位置（但不能回退）
			if idx+1 > start {
				start = idx + 1
			}
		}
		// 更新当前字符的位置
		charIndex[s[end]] = end
		// 更新最大长度
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}

func main() {
	fmt.Print(lengthOfLongestSubstring("abcabcbb"))
}

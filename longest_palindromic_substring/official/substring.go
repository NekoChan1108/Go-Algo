// 官方解动态规划
package main

import "fmt"

func longestPalindrome(s string) string {
	//记录最终返回最长字串的两端下标
	idx1, idx2 := 0, 0
	//记录状态
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		// 当前dp含义是以i为后缀的是否回文
		//单个字符就是回文串记为true
		dp[i] = true
		for j := 0; j < i; j++ {
			//s[j:i+1]为回文串要保证s[j+1:i]为回文串并且两端字符要相等s[i]==s[j]
			// fmt.Println(dp[j+1])
			dp[j] = dp[j+1] && s[i] == s[j]
			// fmt.Println(dp[j])
			if dp[j] && i-j > idx2-idx1 {
				//如果当前回文子串长度大于之前的回文子串长度就更新
				idx1 = j
				idx2 = i
			}
		}
	}
	fmt.Println(dp)
	return s[idx1 : idx2+1]
}

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
	// fmt.Println(Judge("cbb"))
}

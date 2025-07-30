// 自己写的返回最长回文子串
package main

import "fmt"

//判断是不是回文串
func Judge(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	strSlice := make([]string, 0)
	maxStr := ""
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		strSlice = append(strSlice, string(s[i]))
		for j := i + 1; j < len(s); j++ {
			if Judge(s[i:j+1]) && len(s[i:j+1]) > len(maxStr) {
				maxStr = s[i : j+1]
				strSlice = append(strSlice, s[i:j+1])
			}
		}
	}
	fmt.Println(strSlice)
	if len(maxStr) == 0 {
		maxStr = strSlice[0]
	}
	// for i := 0; i < len(strSlice); i++ {
	// 	if len(strSlice[i]) > len(maxStr) {
	// 		maxStr = strSlice[i]
	// 	}
	// }
	return maxStr
}

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
	// fmt.Println(Judge("cbb"))
}

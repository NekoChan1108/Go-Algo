// 有效的括号(自己写的)
package main

import "fmt"

// str:="{[()]}"
// 123 91 40 41 93 125
func isValid(s string) bool {
	//奇数无效
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	//非正向直接返回错误
	if s[0] == '}' || s[0] == ')' || s[0] == ']' {
		return false
	}
	//入栈与比较
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		if v > stack[len(stack)-1] && (v+stack[len(stack)-1] == rune(81) || v+stack[len(stack)-1] == rune(184) || v+stack[len(stack)-1] == rune(248)) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	// if len(stack) == 0 {
	// 	return true
	// }
	// return false
	return len(stack)==0
}
func main() {
	fmt.Println(isValid("([]{})"))
}

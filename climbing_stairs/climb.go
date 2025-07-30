// 自己写的(爬梯子)
package main

import (
	"fmt"
)

//每次只能爬1步或者2步
//也就是最后爬完n个阶梯只能迈1步或者2步
//所以总共的爬法就是到只迈1步的爬法+到只迈2步的爬法
//通项F(n)=F(n-1)+F(n-2)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	//p，q存前两个的解法 res存最终解法
	var p, q, res int
	p = 1
	q = 2
	for i := 3; i <= n; i++ {
		res = p + q
		p = q
		q = res
	}
	return res
}

func main() {
	fmt.Print(climbStairs(5))
}

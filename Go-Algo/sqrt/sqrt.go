// 自己写的(开方 超时)
package main

import "fmt"

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	val := 10
	tmp := x
	var cnt int
	for tmp/10 > 0 {
		cnt++
		tmp /= 10
	}
	for i := 1; i <= cnt-1; i++ {
		val *= 10
	}
	var sqrt int
	for ; val >= 1; val-- {
		if val*val <= x {
			sqrt = val
			return sqrt
		}

	}
	return 0
}

func main() {
	fmt.Print(mySqrt(1085817232))
}

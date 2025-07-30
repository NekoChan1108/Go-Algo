// 官方题解(二分)
package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (r + l) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
func main() {
	fmt.Print(mySqrt(1085817232))
}

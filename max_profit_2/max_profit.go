// 自己写的最大利润2
package main

import "fmt"

//只要前一天买入后一天卖出利润是大于零就可以
func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) && prices[i+1]-prices[i] > 0 {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(s1))
	fmt.Println(maxProfit(s2))
}

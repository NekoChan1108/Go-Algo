// 买股票的最佳时机(自己写的)
package main

import (
	"fmt"
	"math"
	"sort"
)

func Reverse(prices []int) {
	for i := 0; i < len(prices)/2; i++ {
		tmp := prices[i]
		prices[i] = prices[len(prices)-i-1]
		prices[len(prices)-i-1] = tmp
	}
}

// 双重for会超时
func maxProfit(prices []int) int {
	//如果是倒序股票价格则直接返回0
	Reverse(prices)
	if sort.IntsAreSorted(prices) {
		return 0
	}
	Reverse(prices)
	//记录最大利润
	var max = math.MinInt
	//记录前i天最小的买入价格
	var min = math.MaxInt
	for i := 0; i < len(prices); i++ {
		//更新最小买入价格
		if prices[i] < min {
			min = prices[i]
		}
		//更新最大利润
		if i+1 < len(prices) && prices[i+1]-min > max {
			max = prices[i+1] - min
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4, 100, 0, 2}))
}

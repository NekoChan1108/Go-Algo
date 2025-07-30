// 官方解 动态规划
package main

import "fmt"

//持股上分两种状态 一类是持股买入 另一类是无股卖出
//时间上分两种状态 一类是当天 另一类是前一天
func maxProfit(prices []int) int {
	n := len(prices)
	//行表示每一天 列表示持股或者未持股
	dp := make([][2]int, n)
	//第一天未持股利润就是0
	dp[0][0] = 0
	//第一天持股利润即买入也就是-prices[0]
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		//当天未持股最大利润分两种一种就是当天股票价格比前一天低继续不买入
		//第二就是当天股票高前一天持股当天卖出获取当天股票价格利润
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		//当天持股最大利润分两种一种就是当天股票价格比前一天低继续不买入保持前一天所持股
		//第二就是当天股票高前一天未持股当天买入股票
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 1, 5, 3, 6, 4, 10}
	fmt.Println(maxProfit(s1))
	fmt.Println(maxProfit(s2))
}

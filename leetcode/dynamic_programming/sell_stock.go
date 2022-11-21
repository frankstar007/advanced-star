/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/19
 * @contact frankstarye@tencent.com
 * @desc 买卖股票的最佳时机 id:121
 **/

package dynamic_programming

//maxProfit 买入股票的最佳时机 普通方法 容易超时
func maxProfit(prices []int) int {
	max := 0
	//i 代表 第i天买入
	for i := 0; i < len(prices) - 1; i ++ {
		// j 代表 第j天买入
		for j := 1; j < len(prices); j ++ {
			if i >= j {
				continue
			}
			profit := prices[j] - prices[i]
			if profit > max {
				max = profit
			}
		}
	}

	return max
}

func dpMaxProfit(prices []int) int {
	//dp[i] 代表 第i天内的最低价
	max := 0
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	for i := 1; i < len(prices); i ++ {
		if dp[i-1] > prices[i] {
			dp[i] = prices[i]
		}  else {
			dp[i] = dp[i-1]
		}
		if prices[i] - dp[i] > max {
			max = prices[i] - dp[i]
		}
	}
	return max
}
/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 * @desc climbing stairs' id : 70
**/

package easy


func climbStairs(n int) int {
	dp := make([]int, n + 1)
	//其中dp[i] 代表 第i层有几种走法 每层可以看做是走到第i - 1层再走一步  或者  第 i -2 层 再走2步
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i ++ {
		dp[i] = dp[i - 1] + dp[ i - 2]
	}
	return dp[n]


}
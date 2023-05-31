/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/3
 * @contact frankstarye@tencent.com
 * @desc 最长回文子串
 **/

package medium

//longestPalindrome 中心扩散算法
func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return s
	}

	strLen := len(s)
	left, right := 0, 0
	initLen, maxLen, maxStart := 1, 0, 0
	for i := 0; i < strLen; i ++ {
		left, right = i - 1, i + 1

		for left >= 0 && s[i] == s[left] {
			initLen ++
			left --
		}

		for right < strLen && s[i] == s[right] {
			initLen ++
			right ++
		}

		for left >= 0 && right < strLen && s[left] == s[right] {
			initLen += 2
			left --
			right ++
		}

		if initLen > maxLen {
			maxLen = initLen
			maxStart = left
		}
		initLen = 1

	}
	return s[maxStart + 1:maxStart + maxLen + 1]

}

//longestPalindromic 动态规划
func longestPalindromic(s string) string {
	if len(s) <= 0 {
		return s
	}
	dp := make([][]bool, len(s))  // dp[i][j] 代表 i 到 j 是回文字符串
	length := 1
	start, end := 0, 0
	for k := 1; k < len(s); k ++ {
		for j := 0; j < k; j ++ {
			if (s[j] == s[k] && ( k - j <= 2)) || dp[j + 1][k - 1] {
				dp[j][k] = true
				if k - j + 1 > length {
					length = k - j + 1
					start = j
					end = k
				}
			}
		}
	}
	return s[start:end + 1]

}

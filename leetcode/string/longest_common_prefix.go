/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/9
 * @contact frankstarye@tencent.com
 * @desc longest common prefix id:14
**/

package string

import (
	"math"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	min := math.MaxInt
	minIndex := 0
	//先找到最小字符串的位置
	for index, str := range strs {
		length := len(str)
		if length <= min {
			min = length
			minIndex = index
		}

	}
	//用最短字符串作为前缀进行去匹配
	minStr := strs[minIndex]
	flag := false
	for len(minStr) > 0 {

		for _, str := range strs {
			if !strings.HasPrefix(str, minStr) {
				if len(minStr) - 1  >= 0 {
					minStr = minStr[:len(minStr) - 1]
					flag = true
					break
				} else {
					return ""
				}
			}
		}
		if flag {
			flag = false
			continue
		}
		return minStr
	}
	return minStr

}
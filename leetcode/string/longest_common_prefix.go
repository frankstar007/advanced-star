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


//分治法
func lcp(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	return recursionLCP(strs, 0, len(strs) - 1)
}

func recursionLCP(strs []string, start int, end int) string {
	if start == end {
		return strs[start]
	}else {
		mid := (end - start) / 2 + start
		left := recursionLCP(strs, start, mid - 1)
		right := recursionLCP(strs, mid + 1, end)
		return commonPrefix(left, right)
	}

}

func commonPrefix(left string, right string) string {
	minLength := len(right)
	if len(left) < len(right) {
		minLength = len(left)
	}
	for i := 0; i < minLength; i ++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:minLength]
}
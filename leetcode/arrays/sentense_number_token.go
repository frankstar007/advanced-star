/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/3
 * @contact frankstarye@tencent.com
 * @desc 检查句子中的数字是否是递增
**/

package arrays

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func areNumbersAscending(s string) bool {
	tokens := make([]string, 0)
	tokens = append(tokens, strings.Split(s, " ")...)
	pre := math.MinInt
	for _, token := range tokens {
		if strings.ContainsAny(token, "0123456789") {
			//如果是数字token
			number,_ := strconv.Atoi(token)
			if number <= pre {
				return false
			}
			pre = number
		}
	}
	return true
}

//sAreNumbersAscending 官方题解
func sAreNumbersAscending(s string) bool {
	pre, i := 0, 0
	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			cur := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				cur = cur*10 + int(s[i]-'0')
				i++
			}
			if cur <= pre {
				return false
			}
			pre = cur
		} else {
			i++
		}
	}
	return true
}


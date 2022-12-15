/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/8
 * @contact frankstarye@tencent.com
 * @desc 回文子串
 **/

package arrays

import "strconv"

func isReverseNumber(number int) bool {
	strNumber := strconv.Itoa(number)
	length := len(strNumber)
	start, end := 0, length - 1
	for start < end {
		if strNumber[start] == strNumber[end] {
			start ++
			end --
		} else {
			return false
		}
	}
	return true
}
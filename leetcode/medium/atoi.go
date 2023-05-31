/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/11
 * @contact frankstarye@tencent.com
 * @desc id : 8
 **/

package medium

func myAtoi(s string) int {
	if len(s) <= 0 {
		return 0
	}
	index := 0
	flag := false
	var number int
	for index < len(s) && s[index] == ' ' {
		index ++
	}
	if index >= len(s) {
		return number
	}
	if s[index] == '-' || s[index] == '+' {
		if s[index] == '-' {
			flag = true
		}
		index ++
	}

	for index < len(s)  {
		if s[index] < '0' || s[index] > '9' {
			if flag {
				return -number
			}
			return number
		}
		gap := int(s[index] - '0')

		//如果是负数越界
		if flag && (number > 214748364 || (number == 214748364 && gap >= 8)) {
			return -2147483648
		}
		//如果是正数越界
		if !flag && (number > 214748364 || (number == 214748364 && gap >= 7)) {
			return 2147483647
		}
		number = number * 10 + gap


		index ++
	}

	if flag {
		return -number
	}
	return number
}

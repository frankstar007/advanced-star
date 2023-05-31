/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/10
 * @contact frankstarye@tencent.com
 * @desc id:7 reverse integer
 **/

package medium

func reverse(x int) int  {
	if x == 0 {
		return x
	}
	var number int
	flag := false
	if x > 0 {
		flag = true
	} else {
		x = -x
	}

	for x > 0 {
		left := x % 10
		x = x / 10
		if flag && (number > 214748364  || (number == 214748364 && left >= 7)) {
			return 0
		}
		if !flag && (number > 214748364 || (number == 214748364 && left >= 8)) {
			return 0
		}
		number = number * 10 + left

	}

	if !flag {
		return -number
	}


	return number
}
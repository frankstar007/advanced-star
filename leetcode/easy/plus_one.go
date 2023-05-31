/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 * @desc plus one id : 66
 **/

package easy

func plusOne(digits []int) []int {
	flag  := false
	last := digits[len(digits) - 1]
	k := last + 1
	if k >= 10 {
		flag = true
		digits[len(digits) - 1] = k % 10
	} else {
		digits[len(digits) - 1] = k
	}


	for i := len(digits) - 2; i >= 0 ; i-- {
		var k int
		k = digits[i]
		if flag {
			k = k + 1
			if k >= 10 {
				flag = true
				k = k % 10
			} else {
				flag = false
			}
		}
		digits[i] = k
	}

	//如果最后一次相加后 大于10  说明还需要增加一个
	if flag {


		digits = append(digits, 1)
		swap(digits, 0, len(digits) - 1)
	}
	return digits
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}


func perfectPlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i ++ {
		digits[i] ++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append(digits, 1)
	swap(digits, 0, len(digits) - 1)
	return digits
}
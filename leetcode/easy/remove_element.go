/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/17
 * @contact frankstarye@tencent.com
 * @desc remove element id : 27
 **/

package easy

import "fmt"

//removeElement
func removeElement(nums []int, val int) int {
	if len(nums) <= 0 {
		return 0
	}

	l, r := 0, len(nums) - 1

	for l <= r {

		//从右至左找到第一个不等于val的值
		for l <= r && nums[r] == val {
			r --
		}

		//从左至右找到等于val的值
		for l <= r && nums[l] != val {
			l ++
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}

	}
	fmt.Println(fmt.Sprintf("num:%v", nums))
	return l
}

//rmElement 快慢指针
func rmElement(nums []int, val int) int {
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow ++
		}
		fast ++
	}
	return slow
}
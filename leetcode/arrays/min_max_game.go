/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/15
 * @contact frankstarye@tencent.com
 * @desc minMaxGame id:2293
 **/

package arrays

import "fmt"

func minMaxGame(nums []int) int {

	for len(nums) > 1 {
		flag := true
		newNums := make([]int, 0)
		for i := 0; i < len(nums) - 1 && i % 2 == 0; i+= 2 {
			var item int
			if flag  {
				item = getMin(nums, i)
				flag = false
			} else {
				item = getMax(nums, i)
				flag = true
			}
			newNums = append(newNums, item)
		}
		nums = newNums
		fmt.Println(fmt.Sprintf("newNums:%v", nums))
	}
	return nums[0]


}

func getMax(nums []int, i int) int {
	if nums[i] > nums[i + 1] {
		return nums[i]
	}
	return nums[i + 1]
}

func getMin(nums []int, i int) int {
	if nums[i] < nums[i + 1] {
		return nums[i]
	}
	return nums[i + 1]
}
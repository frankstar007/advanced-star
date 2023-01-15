/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/10
 * @contact frankstarye@tencent.com
 * @desc 3sum id : 15
 **/

package sort

import (
	"fmt"
)

func threeSum(nums []int) [][]int {

	//quick sort
	quickSort(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i ++  {

		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		//找到第一个不相等数
		l, r := i + 1, len(nums) - 1
		for l < r {
			target := nums[i] + nums[l] + nums[r]
			if target == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				//此处也需要做判断
				for l < r && nums[l + 1] == nums[l] {
					l ++
				}

				for l < r && nums[r] == nums[r - 1] {
					r --
				}
				l ++
				r --

			} else if target < 0 {
				l ++
			} else {
				r --
			}


		}
	}
	fmt.Println(fmt.Sprintf("three sum%v", result))
	return result
}

//quickSort 快速排序
func quickSort(nums []int) {
	fmt.Println(fmt.Sprintf("before quick sort:%v", nums))
	sort(nums, 0, len(nums) - 1)
	fmt.Println(fmt.Sprintf("after quick sort:%v", nums))

}

func sort(nums []int, start, end int) {
	if start < end {
		l, r := start, end
		pivot := nums[l]
		for l < r {

			//从右至左找到比基准数大的
			for l < r && nums[r] > pivot {
				r --
			}
			if l < r {
				nums[l] =  nums[r]
				l ++
			}
			//从左至右找到比基准值小的
			for l < r && nums[l] < pivot {
				l ++
			}
			if l < r {
				nums[r] = nums[l]
				r --
			}
		}
		nums[l] = pivot
		sort(nums, start, l - 1)
		sort(nums, l + 1, end)

	}
}
/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/15
 * @contact frankstarye@tencent.com
 * @desc merge intervals id :56
 **/

package sort

import (
	"fmt"
	common "github.com/frankstar007/advanced-star/algo/sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	result := make([][]int, 0)
	var i int
	//根据索引0 进行选择排序
	for i = 0; i < len(intervals) - 1; i ++ {
		minIndex := i
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] < intervals[minIndex][0] {
				minIndex = j
			}
		}

		if minIndex != i {
			intervals[i], intervals[minIndex] = intervals[minIndex], intervals[i]
		}
	}
	//fmt.Println(fmt.Sprintf("sort by zero index:%v", intervals))

	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i ++ {
		if intervals[i][0] > end {
			//分开
			item := []int{start, end}
			result = append(result, item)

			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			end = common.MaxInt(end, intervals[i][1])
		}
	}
	item := []int{start, end}
	result = append(result, item)

	//fmt.Println(fmt.Sprintf("result:%v", result))

	return result


}

func sequence(nums []int) {
	fmt.Println(fmt.Sprintf("before sort:%v", nums))
	qSort(nums, 0, len(nums) - 1)
	fmt.Println(fmt.Sprintf("after sort:%v", nums))

}

//qSort 快速排序
func qSort(nums []int, start, end int)   {

	if start < end {
		l , r := start, end
		base := nums[l]

		for l < r {

			//从右至左 找到一个比base还要小的额数
			for l < r && nums[r] > base {
				r--
			}

			if l < r {
				nums[l] = nums[r]
				nums[r] = nums[r - 1]
				l ++
			}

			//从左至右 找到一个比base要大的数
			for l < r && nums[l] < base {
				l ++
			}
			if l < r {
				nums[r] = nums[l]
				nums[l] = nums[l + 1]
				r --
			}


		}
		nums[l] = base
		qSort(nums, start, l - 1)
		qSort(nums, l + 1, end)
	}

}


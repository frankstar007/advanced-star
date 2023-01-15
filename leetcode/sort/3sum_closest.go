/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/14
 * @contact frankstarye@tencent.com
 * @desc 3sum the closest id : 16
 **/

package sort

import (
	"fmt"
	"math"
)

//threeSumClosest 最近的三数之和
func threeSumClosest(nums []int, target int) int {

	//先排序
	quickSort(nums)
	var sum int
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			gap, flag := abs(tmp, target)
			if gap <= min {
				min = gap
				sum = tmp
			}
			if flag {
				r --
			} else {
				l ++
			}

		}

	}
	fmt.Println(fmt.Sprintf("min:%d", min))
	return sum
}

func abs(i, j int) (int, bool) {
	if i > j {
		return i - j, true
	}
	return j - i, false
}

/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/10
 * @contact frankstarye@tencent.com
 * @desc remove duplicates from sorted array
**/

package arrays

import "fmt"

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			nums[slow + 1] = nums[fast]
			slow ++
			fast ++
		}
	}
	fmt.Println(fmt.Sprintf("nums:%v", nums))
	return slow + 1
}


func  removeDuplicate(nums []int) int {
	start, end := 0, len(nums) - 1
	max := nums[len(nums) -1 ]
	for start < end {
		var j int
		for j = start + 1; j < len(nums) && nums[start] == nums[j]; j++ {
		}
		//第一个不相等的地方
		index := j
		if index > start {
			//进行移动
			gap := index - start - 1
			for k := index; k < len(nums); k ++ {
				nums[k-gap] = nums[k]
			}
			start ++
			end -= gap
			for m := 0; m < gap ; m ++ {
				nums[len(nums) - 1 -m] = max - 1
			}
		}

	}
	fmt.Println(fmt.Sprintf("nums:%v", nums))
	result := 1
	begin := 0
	for z := begin + 1; z < len(nums); z ++  {
		if nums[z] > nums[z - 1] {
			result ++
		}
	}
	return result
}

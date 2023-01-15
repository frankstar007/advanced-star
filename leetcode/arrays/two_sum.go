/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/8
 * @contact frankstarye@tencent.com
 * @desc two sum id :1
 **/

package arrays

import "fmt"

func twoSum(nums []int, target int) []int {

	recordMap := make(map[int]int)
	for i, num := range nums {
		left := target - num
		recordMap[left] = i
	}
	result := make([]int, 0)
	uinque := make(map[int]bool)
	for i, num := range nums {
		if _, ok := recordMap[num]; ok {
			if _, ok := uinque[i]; !ok {
				if i != recordMap[num] {
					result = append(result, recordMap[num])
					result = append(result, i)
					uinque[i] = true
					uinque[recordMap[num]] = true
				}
			}

		}
	}
	fmt.Println(fmt.Sprintf("result:%v", result))
	return result



}

func standardTwoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for i, num := range nums {
		if p, ok := record[target - num]; ok {
			return []int{i, p}
		}
		record[num] = i
	}
	return nil
}
/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/22
 * @contact frankstarye@tencent.com
 * @desc 数组中的第K个最大元素 id:215
 **/

package arrays

import "fmt"

func findKthLargest(nums []int, k int) int {
	//获取数组的长度
	length := len(nums)
	//快速排序
	quickSort(nums, 0, length - 1)
	fmt.Println(fmt.Sprintf("%v", nums))
	return nums[k]
}

func quickSort(nums []int, start, end int)  {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot - 1)
		quickSort(nums, pivot + 1, end)
	}
}

func partition(nums []int,start int, end int) int{
	pivot := start
	index := pivot + 1
	for i := index; i <= end; i ++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, pivot)
			index += 1
		}
	}
	swap(nums, pivot, index - 1)
	return index - 1
}

func swap(nums []int, i,j int) {
	nums[i], nums[j] = nums[j], nums[i]
}





/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/4
 * @contact frankstarye@tencent.com
 * @desc search insert position ID：35
 **/

package binary_search

//searchInsert 二分查找有序数组的插入位置
func searchInsert(nums []int, target int) int {

	length := len(nums)
	start, end := 0, length - 1

	for start <= end {
		mid := ( start + end ) / 2
		if target <= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}


	}
	return start
}
/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/10
 * @contact frankstarye@tencent.com
 **/

package sort

//swap 交换数组的元素
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

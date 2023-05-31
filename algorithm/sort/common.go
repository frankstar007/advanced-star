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

func Swap(arr []interface{}, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func MaxInt(i, j int) int {
	if i > j {
		return i
	}

	return j
}
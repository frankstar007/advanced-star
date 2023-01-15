/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
 * @desc 冒泡排序
 **/

package sort

import "fmt"

//bubbleSort 冒泡排序
func bubbleSort(nums []int) {
	//一共需要比较 n - 1趟 最后一个数据已经排好序了
	flag := 0
	for i := len(nums) - 1; i > 0; i-- {
		//每一趟 把最大的数据沉下去
		for j := 0; j < i; j++ {
			if nums[j] > nums[j + 1] {
				swap(nums, j, j + 1)
				flag = 1
			}
		}
		//如果数据已经有序 无需遍历 跳出
		if flag == 0 {
			break
		}
	}
	fmt.Println(fmt.Sprintf("bubble sort:%v", nums))
}


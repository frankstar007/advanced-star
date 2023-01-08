/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/4
 * @contact frankstarye@tencent.com
 * @desc  Maximum Value at a Given Index in a Bounded Array id ： 1802
**/

package arrays

//maxValue 给定数组和边界的最大值
func maxValue(n int, index int, maxSum int) int {
	//step1 get average num and remainder
	average := maxSum / n
	remainder := maxSum % n

	if remainder == 0 {
		return average
	}

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr[i] = average
	}

	for remainder > 0 {
		//如果是首位 最大值
		if index == 0 {

		}


	}









	return 0
}
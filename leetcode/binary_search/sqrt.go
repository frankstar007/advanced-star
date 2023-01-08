/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/4
 * @contact frankstarye@tencent.com
 * @desc x sqrt id ：69
 **/

package binary_search


func mySqrt(x int) int {
	start, end := 0, x
	for start <= end {
		mid := start + (end - start) / 2
		if mid * mid <= x {
			start = mid + 1
		}else {
			end = mid - 1
		}
	}
	return end
}
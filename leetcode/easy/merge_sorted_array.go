/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 * @desc merge sorted array id:88
**/

package easy

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  []int{

	fmt.Println(fmt.Sprintf("before merge nums1 : %v", nums1))
	l , r, k :=  m - 1, n - 1, m + n - 1
	for k >= 0 {

		for l >= 0 && r >= 0 && nums1[l] > nums2[r]  {
			nums1[k] = nums1[l]
			l --
			k --
		}

		for r >= 0 && l >= 0 && nums2[r] >= nums1[l] {
			nums1[k] = nums2[r]
			r --
			k --
		}


		if l >= 0 && r < 0 {
			nums1[k] = nums1[l]
			k --
			l --
		}
		if r >= 0 && l < 0{
			nums1[k] = nums2[r]
			k --
			r --
		}

	}

	fmt.Println(fmt.Sprintf("after merge nums1 : %v", nums1))
	return nums1
}

/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 * @desc 按奇偶数排序数组 id:922
 **/

package arrays

func sortArrayByParityII(nums []int) []int {
	var i, j int
	length := len(nums)
	for {
		even := 0
		exist := true
		for i = even; i < length; {
			if nums[i] % 2 == 0 {
				i += 2
			} else {
				even = i
				exist = false
				break
			}
		}
		if exist {
			break
		}

		odd := 1
		for j = odd; j < length;{
			if nums[j] % 2 != 0 {
				j += 2
			} else {
				odd = j
				break
			}
		}


		tmp := nums[even]
		nums[even] = nums[odd]
		nums[odd] = tmp


	}
	return nums
}

//sortArrayByParity 官方答案
func sortArrayByParity(nums []int) []int {
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}


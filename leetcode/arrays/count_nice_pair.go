/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/17
 * @contact frankstarye@tencent.com
 * @desc 1814. Count Nice Pairs in an Array
 * a. 0 <= i < j < nums.length
 * b. nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
 * c. rev(123) = 321 rev(120) = 21
 **/

package arrays

import (
	"strconv"
)

func countNicePairs(nums []int) int {

	//对集合中的每个元素进行rev
	rNum := make(map[int]int)
	for _, num := range nums {
		rNum[num] = rev(num)
	}
	count := 0
	var result []int
	for i := 0; i <= len(nums) - 1; i ++ {
		result = append(result, nums[i] - rNum[nums[i]])
	}
	indexMap := make(map[int][]int)
	for i, v := range result {
		if _, ok := indexMap[v]; ok {
			indexMap[v] = append(indexMap[v], i)
		} else {
			indexMap[v] = []int{i}
		}
	}

	for _, v := range indexMap {
		length := len(v)
		if length >= 2 {
			count += (length * (length - 1)) / 2
		}

	}
	//for i := 0 ; i < len(nums) - 1; i ++{
	//	left := nums[i] - rNum[nums[i]]
	//	for j := i + 1; j < len(nums); j ++ {
	//		if  nums[j] - rNum[nums[j]] == left {
	//			fmt.Printf("i:%d, j:%d \n", i, j)
	//			count ++
	//		}
	//	}
	//}

	return count
}


func cntNicePairs(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		rev := 0
		for x := num; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		ans += cnt[num-rev]
		cnt[num-rev]++
	}
	return ans % (1e9 + 7)
}



//rev 数字反转
func rev(x int) int {
	//首先转为字符串
	str := strconv.Itoa(x)
	chars := []rune(str)


	flag := true
	//先剔除最后的0 字符
	for flag {
		if chars[len(chars) - 1] - '0' == 0 {
			chars = chars[:len(chars) - 1]
		} else {
			flag = false
		}
	}

	//从最后一个非0字符开始进行reverse
	sum := int(chars[len(chars) - 1]) - '0'
	for i := len(chars) - 2; i >= 0; i-- {
		num := int(chars[i]) - '0'
		sum = sum * 10  + num
	}
	return sum




}
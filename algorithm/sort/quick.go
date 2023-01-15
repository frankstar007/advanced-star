/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
 * @desc 快速排序
 * @思想： 	a.从数列中找出一个基准值
 			b.将所有比基准值小的数移到基准值前，所有比基准值小的数移到基准值后（相同可以放在任意一边）
			c.该分区退出后，该基准就处于数列的中间位置
			d.递归把基准值前面的子数列和基准值后面的子数列进行排序
 **/

package sort

import "fmt"

//quickSort 快速排序 分区排序
func quickSort(nums []int) {

	fmt.Println(fmt.Sprintf("before quick sort nums:%v", nums))
	qSort(nums, 0, len(nums) - 1)
	fmt.Println(fmt.Sprintf("quick sort num:%v", nums))
}


func qSort(nums []int, l, r int) {
	if l < r {
		pt := pivot(nums, l, r)
		qSort(nums, l, pt - 1)
		qSort(nums, pt + 1, r)
	}
}

func pivot(nums []int, l, r int) int {
	base := l
	substitute := base + 1
	for i := substitute; i <= r; i++ {
		if nums[i] < nums[base] {
			swap(nums, i, substitute)
			substitute ++
		}
	}
	swap(nums, base, substitute - 1)
	return substitute - 1
}



func partition(nums []int, l int, r int) int {

	i , j := l, r
	base := nums[i]
	//从右至左 找到小于基准值小的数
	for i < j {
		for i < j && nums[j] > base {
			j--
		}
		//找到右边第一个比基准值小的数 也有可能找不到所以要判断一下
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i] < base {
			i ++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	//最后把基准值放在中间
	nums[i] = base
	return i

}




func quicksort(nums []int, start, end int) {
	fmt.Println(fmt.Sprintf("before quick sort nums:%v", nums))
	if start < end {
		i, j := start, end
		base := nums[i]

		for i < j {
			//先从右至左 找到小于base的值
			for i < j && nums[j] > base {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}

			//在从左至右 找到大于base的值
			for i < j && nums[i] < base {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}

		}

		nums[i] = base
		quicksort(nums, start, i - 1)
		quicksort(nums, i + 1, end)

	}



	fmt.Println(fmt.Sprintf("quick sort num:%v", nums))
}
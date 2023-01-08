/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/7
 * @contact frankstarye@tencent.com
 * @desc  1658. Minimum Operations to Reduce X to Zero
**/

package arrays

import "github.com/frankstar007/advanced-star/collection"

func minOperations(nums []int, x int) int {
	reminder := 0
	left, right := 0, len(nums) - 1
	count := 0
	for left < right && reminder < x{
		if nums[left] > nums[right] {
			reminder += nums[left]
			left ++
			count++
		} else if nums[left] < nums[right] {
			reminder += nums[right]
			right --
			count++
		} else {
			//如果左右相等 就比较 两个的下一个数
			l, r := left + 1, right - 1

			for l < r {
				if nums[l] > nums[r] {
					left ++
					break
				} else if nums[l] < nums[r] {
					right --
					break
				} else {
					l ++
					r --
				}
			}
			if l >= r {
				right --
			}else {
				left ++
			}
			count++

		}
	}
	if left >= right && reminder != x{
		return -1
	}
	return count


}


type Node struct {
	cur int // 0 代表左 1 代表右
	curIndex int // 索引位置
	sub   int // 替补方向
	subIndex int //替补位置

}

//minOperation 最少操作次数
func minOperation(nums []int, x int) int {
	stack := collection.NewStack()
	left, right := 0, len(nums) - 1
	if nums[left] > x && nums[right] > x {
		return -1
	}
	//如果最左边的数大于x 只能从右边进行操作
	count := 0
	if nums[left] > x {
		for i := right; i > left; i -- {
			if x < 0 {
				return -1
			} else if x > 0 {
				x = x - nums[i]
				count ++
			} else {
				return count
			}
		}
	}
	if nums[right] > x {
		for i := left; i < right; i ++ {
			if x < 0 {
				return -1
			} else if x > 0 {
				x = x - nums[i]
				count ++
			} else {
				return count
			}
		}
	}
	//走到这一步 左右两边都同时小于 x
	sum := 0
	for left <= right {
		if sum > x {
			//进行pop
			if stack.Empty() {
				return -1
			}
			//回溯
			node := stack.Pop().(Node)
			sum -= nums[node.curIndex]
			cur := Node{
				cur: node.sub,
				curIndex: node.subIndex,
				sub: node.cur,
				subIndex: node.curIndex,
			}
			stack.Push(cur)
			sum += nums[cur.curIndex]
			if cur.cur == 0 {
				left ++
			} else {
				right --
			}

		} else if sum == x {
			return stack.Len()
		} else  {
			if nums[left] > nums[right]  {
				//如果左边大于右边
				node := Node{
					cur: 0,
					curIndex: left,
					sub: 1,
					subIndex: right,
				}
				stack.Push(node)
				sum += nums[node.curIndex]
				left ++
			} else {
				node := Node{
					cur: 1,
					curIndex: right,
					sub: 0,
					subIndex: left,
				}
				stack.Push(node)
				sum += nums[node.curIndex]
				right --
			}

		}
	}

	return -1
}
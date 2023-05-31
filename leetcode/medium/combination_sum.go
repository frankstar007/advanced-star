/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/20
 * @contact frankstarye@tencent.com
 * @desc id:39 组合数
 **/

package medium

import (
	"github.com/frankstar007/advanced-star/collection"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if candidates[0] > target {
		return [][]int{}
	}

	var end int
	for index, c := range candidates {
		if c >= target {
			end = index
		}
	}

	//元素在这个范围内

	answers := make([][]int, 0)
	stack := collection.NewStack()
	var sum int
	var ans []int
	for end >= 0 {
		stack.Push(end)
		sum += candidates[end]
		if sum == target {
			var last int
			for !stack.Empty() {
				last = stack.Top().(int)
				ans = append(ans, candidates[stack.Pop().(int)])
			}
			answers = append(answers, ans)
			//置为0
			sum = 0
			end = last - 1
		} else if sum < target {
			continue
		} else if sum > target {
			stack.Pop()
			end --
		}
	}

	return answers

}

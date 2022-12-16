/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 * @desc 103 二叉树的锯齿形层序遍历
 **/

package binary_tree

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
)




//zigzagLevelOrder 锯齿形遍历二叉树
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)

	dummy := root
	sk := collection.NewQueue()
	sk.Push(dummy)
	odd := 0
	for !sk.Empty() {
		length := sk.Len()
		item := make([]int, 0)
		for i := 0; i < length; i++ {
			tmp := sk.Pop()
			node :=  tmp.(*TreeNode)
			item = append(item, node.Val)
			if node.Left != nil {
				sk.Push(node.Left)
			}
			if node.Right != nil {
				sk.Push(node.Right)
			}

		}
		if odd % 2 == 0 {
			result = append(result, item)
		} else {
			reverse(item)
			result = append(result, item)
		}
		odd ++
		fmt.Println(fmt.Sprintf("item:%v", item))
	}
	return result
}

func reverse(item []int) {
	start, end := 0, len(item) - 1
	for start < end {
		swap(item, start, end)
		start ++
		end --
	}
}

func swap(item []int, start int, end int) {
	item[start] , item[end] = item[end], item[start]
}




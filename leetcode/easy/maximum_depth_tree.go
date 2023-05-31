/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/24
 * @contact frankstarye@tencent.com
 * @desc maximum depth of binary tree id : 104
 **/

package easy

import "github.com/frankstar007/advanced-star/collection"

//maxDepth 树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxInt(maxDepth(root.Right), maxDepth(root.Left))
}

func maxInt(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func iterateMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := collection.NewQueue()
	queue.Push(root)

	for !queue.Empty() {
		length := queue.Len()
		//每一层
		for i := 0; i < length; i ++ {
			top := queue.Pop().(*TreeNode)
			if top.Left != nil {
				queue.Push(top.Left)
			}
			if top.Right != nil {
				queue.Push(top.Right)
			}
		}
		if length > 0 {
			ans += 1
		}
	}
	return ans
}

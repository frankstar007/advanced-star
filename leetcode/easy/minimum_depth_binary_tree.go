/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/24
 * @contact frankstarye@tencent.com
 * @desc minimum depth of binary tree id ： 111
 **/

package easy

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 +  minInt(minDepth(root.Left), minDepth(root.Right))
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}

/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/24
 * @contact frankstarye@tencent.com
 * @desc balanced binary tree
 **/

package easy

func isBalanced(root *TreeNode) bool {

	if root == nil{
		return true
	}
	leftVal := depth(root.Left)
	rightVal := depth(root.Right)
	return abs(leftVal , rightVal) <= 1 && isBalanced(root.Right) && isBalanced(root.Left)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxInt(depth(root.Right), depth(root.Left))
}

func abs(i , j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/24
 * @contact frankstarye@tencent.com
 * @desc path sum  id :112
 **/

package easy

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val);

}

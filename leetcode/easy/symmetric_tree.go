/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/23
 * @contact frankstarye@tencent.com
 * @desc symmetric tree id : 101
 **/

package easy

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	return p.Val == q.Val && symmetric(p.Right, q.Left) && symmetric(p.Left, q.Right)

}

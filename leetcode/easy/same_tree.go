/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/23
 * @contact frankstarye@tencent.com
 & @desc same tree id : 100
 **/

package easy

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

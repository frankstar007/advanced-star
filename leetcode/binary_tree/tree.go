/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 * @desc 二叉树定义
 **/

package binary_tree

//TreeNode 二叉树的节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Node N叉树的节点定义
type Node struct {
	Val int
	Children []*Node
}
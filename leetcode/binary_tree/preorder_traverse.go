/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/30
 * @contact frankstarye@tencent.com
 * @desc 前序遍历
 **/

package binary_tree

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
)

//preorderTraversal 前序遍历
func preorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)
	recursionPreOrder(root, &res)
	fmt.Println(fmt.Sprintf("recursion preorder arr:%v", res))

	iter := iterationPreorder(root)
	fmt.Println(fmt.Sprintf("iteration preorder arr:%v", iter))
	return res

}

//iterationPreorder 迭代前序遍历二叉树
func iterationPreorder(root *TreeNode) []int {
	res := make([]int, 0)
	if root  == nil {
		return res
	}
	stack := collection.NewStack()

	for root != nil || !stack.Empty() {
		for root != nil {
			res = append(res, root.Val)
			stack.Push(root)
			root = root.Left
		}
		cur := stack.Pop().(*TreeNode)
		root = cur.Right
	}
	return res
}

//recursionPreOrder 递归前序遍历二叉树
func recursionPreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	recursionPreOrder(root.Left, res)
	recursionPreOrder(root.Right, res)
}


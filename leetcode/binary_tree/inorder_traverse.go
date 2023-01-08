/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/30
 * @contact frankstarye@tencent.com
 * @desc 中序遍历
**/

package binary_tree

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
)

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderRecursion(root, &res)
	fmt.Println(fmt.Sprintf("%v", res))
	return res
}

//inorderRecursion 递归中序遍历 二叉树
func inorderRecursion(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorderRecursion(node.Left, res)
	*res = append(*res, node.Val)
	inorderRecursion(node.Right, res)
}

//inorderIteration 迭代中序遍历 二叉树
func inorderIteration(node *TreeNode) []int {
	//如果某个节点为空 直接返回
	res := make([]int, 0)
	if node == nil {
		return res
	}
	stack := collection.NewStack()
	for node != nil || !stack.Empty() {
		//一直把左节点放在栈里
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		//如果 左节点为空 则进行pop
		tmp := stack.Pop().(*TreeNode)
		res = append(res, tmp.Val)

		node = tmp.Right

	}



	fmt.Println(res)
	return res
}
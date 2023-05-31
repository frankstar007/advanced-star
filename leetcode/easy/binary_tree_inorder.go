/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/23
 * @contact frankstarye@tencent.com
 * @desc binary tree inorder traversal id : 94
 **/

package easy

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	inorder(root, &res)
	fmt.Println(fmt.Sprintf("recursive inorder binary tree:%v", res))
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		inorder(root.Right, res)
	}

}


func iterateInorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := collection.NewStack()
	res := make([]int, 0)
	for !stack.Empty() || root != nil {

		//一直把左节点放入栈里 找到最左节点
		for  root != nil  {
			stack.Push(root)
			root = root.Left
		}

		//取栈顶元素 放入结果集
		tmp := stack.Pop().(*TreeNode)
		res = append(res, tmp.Val)
		if tmp.Right != nil {
			root = tmp.Right
		}

	}
	fmt.Println(fmt.Sprintf("iterarte inorder binary tree:%v", res))
	return res

}


type ColorNode struct {
	Visit bool
	Node *TreeNode
}


func colorInorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := collection.NewStack()
	cn := &ColorNode{
		Visit: false,
		Node: root,
	}
	stack.Push(cn)
	res := make([]int, 0)
	for !stack.Empty() {
		tmp := stack.Pop().(*ColorNode)
		if tmp.Node == nil {
			continue
		}
		if tmp.Visit {
			res = append(res, tmp.Node.Val)
		} else {
			rcn := &ColorNode{
				Visit: false,
				Node: tmp.Node.Right,
			}
			stack.Push(rcn)
			mcn := &ColorNode{
				Visit: true,
				Node: tmp.Node,
			}
			stack.Push(mcn)

			lcn := &ColorNode{
				Visit: false,
				Node: tmp.Node.Left,
			}
			stack.Push(lcn)
		}

	}
	fmt.Println(fmt.Sprintf("color inorder binary tree:%v", res))

	return res
}

/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/2
 * @contact frankstarye@tencent.com
 * @desc 验证是否是二叉搜索树
 **/

package binary_tree

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
	"math"
)

func isValidBST(root *TreeNode) bool {
	return recursionValidBST(root, math.MinInt, math.MaxInt)
}


func recursionValidBST(root *TreeNode, min, max int) bool {
	if root  == nil {
		return true
	}

	if root.Val <= min || root.Val >= max{
		return false
	}
	return recursionValidBST(root.Left, min, root.Val) &&
			recursionValidBST(root.Right, root.Val, max)
}

//iterateValidBST 迭代方式验证二叉搜索树
func iterateValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pre := math.MinInt
	stack := collection.NewStack()
	for !stack.Empty() || root != nil {

		//左节点放在栈里
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		root = stack.Pop().(*TreeNode)
		fmt.Println(root.Val)
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		root = root.Right
	}
	return true
}
/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/30
 * @contact frankstarye@tencent.com
 **/

package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, deletes []int) []*TreeNode {

	rs := make([]*TreeNode, 0)
	del(root, true, deletes, &rs)
	return rs

}

func del(node *TreeNode, isRoot bool, deletes []int, rs *[]*TreeNode) *TreeNode {

	if node == nil {
		return nil
	}

	if include(node.Val, deletes) {
		node.Left = del(node.Left, true, deletes, rs)
		node.Right = del(node.Right, true, deletes, rs)
		return nil
	} else {
		if isRoot {
			*rs = append(*rs, node)
		}
		return node
	}

}

func include(ele int, eles []int) bool {
	for _, e := range eles {
		if e == ele {
			return true
		}
	}
	return false
}

func dfs(root *TreeNode, deletes []int) [][]int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = push(stack, root)

	var node *TreeNode
	rs := make([][]int, 0)
	result := make([]int, 0)
	for len(stack) > 0 {
		node, stack = pop(stack)
		if include(node.Val, deletes) {
			result = append(result, -1)
			rs = append(rs, result)
			result = []int{}
		} else {
			result = append(result, node.Val)
		}
		if node.Right != nil {
			stack = push(stack, node.Right)
		}
		if node.Left != nil {
			stack = push(stack, node.Left)
		}
	}

	rs = append(rs, result)

	return rs
}

func push(stack []*TreeNode, ele *TreeNode) []*TreeNode {
	stack = append(stack, ele)
	return stack
}

func pop(stack []*TreeNode) (*TreeNode, []*TreeNode) {
	ele := stack[len(stack)-1]
	if len(stack) == 1 {
		return ele, []*TreeNode{}
	}
	return ele, stack[:len(stack)-1]
}




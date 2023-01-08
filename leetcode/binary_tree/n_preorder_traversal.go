/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/30
 * @contact frankstarye@tencent.com
 * @desc N叉树的前序遍历
**/

package binary_tree

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
)

//preorder N叉树的前序遍历
func preorder(root *Node) []int {
	res := make([]int, 0)
	recursionPreorder(root, &res)
	fmt.Println(fmt.Sprintf("recursion preorder :%v", res))


	iter := iteratePreorder(root)
	fmt.Println(fmt.Sprintf("recursion preorder :%v", iter))


	op := optimizePreorder(root)
	fmt.Println(fmt.Sprintf("optimize preorder :%v", op))

	return res
}


func recursionPreorder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, child := range root.Children {
		recursionPreorder(child, res)
	}

}

//iteratePreorder 迭代前序遍历N叉树
func iteratePreorder(root *Node) []int {
	res := make([]int,0)
	if root == nil {
		return res
	}
	stack := collection.NewStack()
	visit := make(map[*Node]int)
	for !stack.Empty() || root != nil {

		for root != nil {
			res = append(res, root.Val)
			stack.Push(root)
			if len(root.Children) > 0 {
				visit[root] = 1
				root = root.Children[0]
			} else {
				break
			}
		}
		cur := stack.Top().(*Node)
		index := visit[cur]
		if index < len(cur.Children) {
			visit[cur] = index + 1
			root = cur.Children[index]
		} else {
			stack.Pop()
			delete(visit, cur)
			root = nil
		}

	}
	return res
}

//optimizePreorder 优化前序遍历N叉树
func optimizePreorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := collection.NewStack()
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Pop().(*Node)
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0 ; i-- {
			stack.Push(node.Children[i])
		}
	}
	return ans
}
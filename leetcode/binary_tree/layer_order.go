/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 * @desc 层序遍历
 **/

package binary_tree

import (
	"fmt"
	"github.com/frankstar007/advanced-star/collection"
)

//layerOrder 层序遍历
func layerOrder(root *TreeNode) [][]int {

	//如果树为空
	if root == nil {
		return [][]int{}
	}

	sk := collection.NewQueue()
	sk.Push(root)
	result := make([][]int, 0)
	for !sk.Empty() {
		length := sk.Len()
		layer := make([]int, 0)
		for i := 0; i < length; i++ {
			node := sk.Pop().(*TreeNode)
			layer = append(layer, node.Val)
			if node.Left != nil {
				sk.Push(node.Left)
			}
			if node.Right != nil {
				sk.Push(node.Right)
			}
		}
		result = append(result, layer)
	}
	fmt.Println(fmt.Sprintf("%v", result))
	return result
}
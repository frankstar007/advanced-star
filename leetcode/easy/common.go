/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 **/

package easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Equal(head *ListNode) bool {
	k := l
	for k != nil && head != nil && k.Val == head.Val {
		k = k.Next
		head = head.Next
	}
	return k == nil && head == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Equal(o *TreeNode) bool  {
	this := t
	thisRes := inorderTraversal(this)
	oRes :=  inorderTraversal(o)
	if len(thisRes) != len(oRes) {
		return false
	}
	for i := range thisRes {
		if thisRes[i] != oRes[i] {
			return false
		}
		fmt.Print(thisRes[i])
	}
	fmt.Println()
	return true


}


/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/19
 * @contact frankstarye@tencent.com
 * @desc 链表节点定义
 **/

package link_list

import (
	"fmt"
	"strings"
)

//ListNode 链表节点
type ListNode struct {
	Val int
	Next *ListNode
}


func (L *ListNode) ToString() string {
	sb := strings.Builder{}
	p := L
	for p != nil {
		sb.WriteString(fmt.Sprint(p.Val))
		sb.WriteString(",")
		p = p.Next
	}
	return strings.TrimRight(sb.String(), ",")
}

func (L *ListNode) Equal(o *ListNode) bool {
	this := L
	if o == nil {
		return false
	}
	for this != nil && o != nil {
		if this.Val != o.Val {
			return false
		}
		this =  this.Next
		o = o.Next
	}
	if o != nil && this != nil{
		return false
	}
	return  true
}
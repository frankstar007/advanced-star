/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 * @desc add two numbers id:2
 **/

package link_list


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//如果l1 链表为空 返回 l2
	if l1 == nil {
		return l2
	}
	//如果l2 链表为空 返回l1
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{
		Val: -1,
	}
	newHead := dummy

	flag := false
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val
		if flag {
			tmp += 1
		}
		if tmp >= 10 {
			flag = true
		} else {
			flag = false
		}
		val := (tmp) % 10
		node := &ListNode{
			Val: val,
		}
		dummy.Next = node
		l1 = l1.Next
		l2 = l2.Next
		dummy = dummy.Next
	}

	for l1 != nil {
		tmp := l1.Val
		if flag {
			tmp += 1
		}
		if tmp >= 10 {
			flag = true
		} else {
			flag = false
		}
		node := &ListNode{
			Val: tmp % 10,
		}
		dummy.Next = node
		l1 = l1.Next
		dummy = dummy.Next
	}


	for l2 != nil {
		tmp := l2.Val
		if flag {
			tmp += 1
		}
		if tmp >= 10 {
			flag = true
		} else {
			flag = false
		}
		node := &ListNode{
			Val: tmp % 10,
		}
		dummy.Next = node
		l2 = l2.Next
		dummy = dummy.Next
	}

	if flag {
		node := &ListNode{
			Val: 1,
		}
		dummy.Next = node
	}

	return newHead.Next
}


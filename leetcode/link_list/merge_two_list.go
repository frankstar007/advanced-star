/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/26
 * @contact frankstarye@tencent.com
 * @desc merge two sorted list id:21
 **/

package link_list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	p := list1
	q := list2
	newHead := &ListNode{
		Val: -1,
	}
	k := newHead
	for p != nil && q != nil {
		if p.Val <= q.Val {
			k.Next = p
			p = p.Next
		}else {
			k.Next = q
			q = q.Next
		}
		k = k.Next
	}
	for p != nil {
		k.Next = p
		p = p.Next
		k = k.Next
	}
	for q != nil {
		k.Next = q
		q = q.Next
		k = k.Next
	}

	return newHead.Next
}


func recursionMergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = recursionMergeTwoList(list1.Next, list2)
		return list1
	} else {
		list2.Next = recursionMergeTwoList(list1, list2.Next)
		return list2
	}
}
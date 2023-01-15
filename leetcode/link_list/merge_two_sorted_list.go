/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/9
 * @contact frankstarye@tencent.com
 **/

package link_list

func mergeTwoSortedListsWithNew(list1, list2 *ListNode) *ListNode {
	p := list1
	q := list2
	if p == nil {
		return q
	}

	if q == nil {
		return p
	}

	dummy := &ListNode{
		Val: - 1,
	}
	newHead := dummy

	for p != nil && q != nil {
		if p.Val < q.Val {
			dummy.Next = p
			dummy = p
			p = p.Next
		} else {
			dummy.Next = q
			dummy = q
			q = q.Next
		}
	}
	//如果p 还不为空 则继续
	for p != nil {
		dummy.Next = p
		dummy = p
		p = p.Next
	}
	//如果q 还不为空 则继续
	for q != nil {
		dummy.Next = q
		dummy = q
		q = q.Next
	}



	return newHead.Next

}



func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	//先找到小的作为主分支
	var m *ListNode
	var s *ListNode
	if list1.Val < list2.Val {
		m = list1
		s = list2
	} else {
		m = list2
		s = list1
	}
	newHead := m

	for m.Next != nil && s != nil{
		if m.Next.Val < s.Val {
			m = m.Next
		} else {
			tmp := s.Next

			s.Next = m.Next
			m.Next = s
			m = m.Next
			s = tmp
		}
	}

	//if m is nil then concat s
	if m.Next == nil {
		m.Next = s
	}
	//if s is nil  then return
	return newHead
}


func recursionMergeTwoSortedList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = recursionMergeTwoSortedList(list1.Next, list2)
		return list1
	} else {
		list2.Next = recursionMergeTwoSortedList(list1, list2.Next)
		return list2
	}

}
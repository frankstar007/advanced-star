/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/19
 * @contact frankstarye@tencent.com
 * @desc 链表反转 ID:206
 **/

package link_list

//reverseList 反转链表节点  非递归方式
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil{
		return head
	}

	var pre *ListNode
	cur := head
	for cur.Next != nil {
		post := cur.Next

		cur.Next = pre
		pre = cur
		cur = post
	}

	return pre
}

//iteratorReverseList 反转链表节点 递归方式
func iteratorReverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	newHead := iteratorReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead


}


